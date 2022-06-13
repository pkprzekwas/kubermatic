//go:build ee

/*
                  Kubermatic Enterprise Read-Only License
                         Version 1.0 ("KERO-1.0”)
                     Copyright © 2021 Kubermatic GmbH

   1.	You may only view, read and display for studying purposes the source
      code of the software licensed under this license, and, to the extent
      explicitly provided under this license, the binary code.
   2.	Any use of the software which exceeds the foregoing right, including,
      without limitation, its execution, compilation, copying, modification
      and distribution, is expressly prohibited.
   3.	THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
      EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
      MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
      IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
      CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
      TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
      SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

   END OF TERMS AND CONDITIONS
*/

package resourcequotas

import (
	"context"
	"encoding/json"
	apiv1 "k8c.io/kubermatic/v2/pkg/api/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"net/http"

	"github.com/gorilla/mux"

	k8cv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/provider"
	utilerrors "k8c.io/kubermatic/v2/pkg/util/errors"
)

// swagger:parameters getResourceQuota
type getResourceQuota struct {
	// in: path
	// required: true
	Name string `json:"name"`
}

// swagger:parameters listResourceQuotas
type listResourceQuotas struct {
	// in: query
	// required: false
	SubjectName string `json:"subject_name,omitempty"`

	// in: query
	// required: false
	SubjectKind string `json:"subject_kind,omitempty"`
}

// swagger:parameters getResourceQuota
type createResourceQuota struct {
	// in: body
	// required: true
	Body struct {
		Subject k8cv1.Subject
		Quota   k8cv1.ResourceDetails
	}
}

type updateResourceQuota struct {
	// in: path
	// required: true
	Name string `json:"name"`

	// in: body
	Body struct {
		CPU     *resource.Quantity `json:"cpu,omitempty"`
		Memory  *resource.Quantity `json:"memory,omitempty"`
		Storage *resource.Quantity `json:"storage,omitempty"`
	}
}

func (m createResourceQuota) Validate() error {
	if m.Body.Subject.Name == "" {
		return utilerrors.NewBadRequest("subject's name cannot be empty")
	}

	if m.Body.Subject.Kind == "" {
		return utilerrors.NewBadRequest("subject's kind cannot be empty")
	}

	return nil
}

func DecodeResourceQuotaReq(r *http.Request) (interface{}, error) {
	var req getResourceQuota

	req.Name = mux.Vars(r)["name"]

	if req.Name == "" {
		return nil, utilerrors.NewBadRequest("`name` cannot be empty")
	}

	return req, nil
}

func DecodeListResourceQuotaReq(r *http.Request) (interface{}, error) {
	var req listResourceQuotas

	req.SubjectName = r.URL.Query().Get("subjectName")
	req.SubjectKind = r.URL.Query().Get("subjectKind")

	return req, nil
}

func DecodeCreateResourceQuotaReq(r *http.Request) (interface{}, error) {
	var req createResourceQuota

	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeUpdateResourceQuotaReq(r *http.Request) (interface{}, error) {
	var req updateResourceQuota

	req.Name = mux.Vars(r)["name"]

	if req.Name == "" {
		return nil, utilerrors.NewBadRequest("`name` cannot be empty")
	}

	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, err
	}

	return req, nil
}

func GetResourceQuota(ctx context.Context, request interface{}, provider provider.ResourceQuotaProvider) (*apiv1.ResourceQuota, error) {
	req, ok := request.(getResourceQuota)
	if !ok {
		return nil, utilerrors.NewBadRequest("invalid request")
	}

	resourceQuota, err := provider.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	resp := &apiv1.ResourceQuota{
		Name:   resourceQuota.Name,
		Spec:   resourceQuota.Spec,
		Status: resourceQuota.Status,
	}

	return resp, nil
}

func ListResourceQuotas(ctx context.Context, request interface{}, provider provider.ResourceQuotaProvider) ([]apiv1.ResourceQuota, error) {
	req, ok := request.(listResourceQuotas)
	if !ok {
		return nil, utilerrors.NewBadRequest("invalid request")
	}

	// TODO: remember to test this part
	labelSet := make(map[string]string)
	if req.SubjectKind != "" {
		labelSet[k8cv1.ResourceQuotaSubjectKindLabelKey] = req.SubjectKind
	}
	if req.SubjectName != "" {
		labelSet[k8cv1.ResourceQuotaSubjectNameLabelKey] = req.SubjectName
	}

	resourceQuotaList, err := provider.List(ctx, labelSet)
	if err != nil {
		return nil, err
	}

	resp := make([]apiv1.ResourceQuota, len(resourceQuotaList.Items))
	for idx, rq := range resourceQuotaList.Items {
		resp[idx] = apiv1.ResourceQuota{
			Name:   rq.Name,
			Spec:   rq.Spec,
			Status: rq.Status,
		}
	}

	return resp, nil
}

func CreateResourceQuota(ctx context.Context, request interface{}, provider provider.ResourceQuotaProvider) error {
	req, ok := request.(createResourceQuota)
	if !ok {
		return utilerrors.NewBadRequest("invalid request")
	}

	if err := req.Validate(); err != nil {
		return utilerrors.NewBadRequest(err.Error())
	}

	if err := provider.Create(ctx, req.Body.Subject, req.Body.Quota); err != nil {
		return err
	}
	return nil
}

func UpdateResourceQuota(ctx context.Context, request interface{}, provider provider.ResourceQuotaProvider) error {
	req, ok := request.(updateResourceQuota)
	if !ok {
		return utilerrors.NewBadRequest("invalid request")
	}

	newQuota := k8cv1.ResourceDetails{
		CPU:     req.Body.CPU,
		Memory:  req.Body.Memory,
		Storage: req.Body.Storage,
	}

	if err := provider.Update(ctx, req.Name, newQuota); err != nil {
		return err
	}
	return nil
}

func DeleteResourceQuota(ctx context.Context, request interface{}, provider provider.ResourceQuotaProvider) error {
	req, ok := request.(getResourceQuota)
	if !ok {
		return utilerrors.NewBadRequest("invalid request")
	}

	if err := provider.Delete(ctx, req.Name); err != nil {
		return err
	}
	return nil
}
