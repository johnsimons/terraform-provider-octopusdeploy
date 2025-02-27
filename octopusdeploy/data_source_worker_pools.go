package octopusdeploy

import (
	"context"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWorkerPools() *schema.Resource {
	return &schema.Resource{
		Description: "Provides information about existing worker pools.",
		ReadContext: dataSourceWorkerPoolsRead,
		Schema:      getWorkerPoolDataSchema(),
	}
}

func dataSourceWorkerPoolsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	query := octopusdeploy.WorkerPoolsQuery{
		IDs:         expandArray(d.Get("ids").([]interface{})),
		PartialName: d.Get("partial_name").(string),
		Skip:        d.Get("skip").(int),
		Take:        d.Get("take").(int),
	}

	client := m.(*octopusdeploy.Client)
	workerPools, err := client.WorkerPools.Get(query)
	if err != nil {
		return diag.FromErr(err)
	}

	flattenedWorkerPools := []interface{}{}
	for _, workerPool := range workerPools.Items {
		workerPoolResource, err := octopusdeploy.ToWorkerPoolResource(workerPool)
		if err != nil {
			return diag.FromErr(err)
		}

		flattenedWorkerPools = append(flattenedWorkerPools, flattenWorkerPool(workerPoolResource))
	}

	d.Set("worker_pools", flattenedWorkerPools)
	d.SetId("Worker Pools " + time.Now().UTC().String())

	return nil
}
