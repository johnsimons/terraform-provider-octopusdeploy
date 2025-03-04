package octopusdeploy

import (
	"context"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const spaceManagersTeamIDPrefix = "teams-spacemanagers-"

func expandSpace(d *schema.ResourceData) *octopusdeploy.Space {
	name := d.Get("name").(string)

	space := octopusdeploy.NewSpace(name)
	space.ID = d.Id()

	if v, ok := d.GetOk("description"); ok {
		space.Description = v.(string)
	}

	if v, ok := d.GetOk("is_default"); ok {
		space.IsDefault = v.(bool)
	}

	if v, ok := d.GetOk("space_managers_team_members"); ok {
		space.SpaceManagersTeamMembers = getSliceFromTerraformTypeList(v)
	}

	if v, ok := d.GetOk("space_managers_teams"); ok {
		space.SpaceManagersTeams = addSpaceManagers(space.GetID(), getSliceFromTerraformTypeList(v))
	}

	if v, ok := d.GetOk("is_task_queue_stopped"); ok {
		space.TaskQueueStopped = v.(bool)
	}

	return space
}

func addSpaceManagers(spaceID string, teamIDs []string) []string {
	var newSlice []string
	if getStringOrEmpty(spaceID) != "" {
		newSlice = append(newSlice, spaceManagersTeamIDPrefix+spaceID)
	}
	for _, v := range teamIDs {
		newSlice = append(newSlice, v)
	}
	return newSlice
}

func flattenSpace(space *octopusdeploy.Space) map[string]interface{} {
	if space == nil {
		return nil
	}

	return map[string]interface{}{
		"description":                 space.Description,
		"id":                          space.GetID(),
		"is_default":                  space.IsDefault,
		"is_task_queue_stopped":       space.TaskQueueStopped,
		"name":                        space.Name,
		"space_managers_team_members": space.SpaceManagersTeamMembers,
		"space_managers_teams":        space.SpaceManagersTeams,
	}
}

func getSpaceDataSchema() map[string]*schema.Schema {
	dataSchema := getSpaceSchema()
	setDataSchema(&dataSchema)

	return map[string]*schema.Schema{
		"id":           getDataSchemaID(),
		"ids":          getQueryIDs(),
		"name":         getQueryName(),
		"partial_name": getQueryPartialName(),
		"skip":         getQuerySkip(),
		"spaces": {
			Computed:    true,
			Description: "A list of spaces that match the filter(s).",
			Elem:        &schema.Resource{Schema: dataSchema},
			Optional:    true,
			Type:        schema.TypeList,
		},
		"take": getQueryTake(),
	}
}

func getSpaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": getDescriptionSchema(),
		"id":          getIDSchema(),
		"is_default": {
			Description: "Specifies if this space is the default space in Octopus.",
			Optional:    true,
			Type:        schema.TypeBool,
		},
		"name": getNameSchema(true),
		"space_managers_team_members": {
			Computed:    true,
			Description: "A list of user IDs designated to be managers of this space.",
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Type:        schema.TypeList,
		},
		"space_managers_teams": {
			Computed:    true,
			Description: "A list of team IDs designated to be managers of this space.",
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Type:        schema.TypeList,
		},
		"is_task_queue_stopped": {
			Description: "Specifies the status of the task queue for this space.",
			Optional:    true,
			Type:        schema.TypeBool,
		},
	}
}

func setSpace(ctx context.Context, d *schema.ResourceData, space *octopusdeploy.Space) error {
	d.Set("description", space.Description)
	d.Set("id", space.GetID())
	d.Set("is_default", space.IsDefault)
	d.Set("name", space.Name)

	if err := d.Set("space_managers_team_members", space.SpaceManagersTeamMembers); err != nil {
		return fmt.Errorf("error setting space_managers_team_members: %s", err)
	}

	if err := d.Set("space_managers_teams", removeSpaceManagers(space.SpaceManagersTeams)); err != nil {
		return fmt.Errorf("error setting space_managers_teams: %s", err)
	}

	d.Set("is_task_queue_stopped", space.TaskQueueStopped)

	return nil
}

func removeSpaceManagers(teamIDs []string) []string {
	if len(teamIDs) == 0 {
		return teamIDs
	}
	var newSlice []string
	for _, v := range teamIDs {
		if !strings.Contains(v, spaceManagersTeamIDPrefix) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
