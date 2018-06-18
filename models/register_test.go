package msaevents

import (
	"encoding/json"
	. "github.com/franela/goblin"
	"testing"
)

func TestEventRegister(t *testing.T) {
	var data EventCreatedUser
	err := json.Unmarshal([]byte(dataEventRegister), &data)
	if err != nil {
		t.Errorf("fail to unmarshal test data: %s", err.Error())
		return
	}

	g := Goblin(t)
	g.Describe("Custom fields", func() {
		labelValues := data.Data.GetCustomFieldsValues("FR")
		g.It("should content 18 custom fields", func() {
			g.Assert(len(data.Data.CustomFields)).Equal(18)
		})
		g.It("should first field be register type", func() {
			g.Assert(data.Data.CustomFields[0].Data.StringValue()).Equal("inscription")
		})
		g.It("should array cast to string", func() {
			g.Assert(data.Data.CustomFields[2].Data.StringValue()).Equal("ab,cd")
		})
		g.It("should int cast to string", func() {
			g.Assert(data.Data.CustomFields[8].Data.StringValue()).Equal("645454545")
		})
		g.It("should have same label values struct than list", func() {
			g.Assert(len(labelValues)).Equal(len(data.Data.CustomFields))
		})
		g.It("should have non empty label", func() {
			g.Assert(labelValues[0].Label).Equal("Type d'inscription")
		})
	})

}

const (
	dataEventRegister = `
{
"config_id":"AAA",
"data": {"id":1267264340256770741,"first_name":"Rémi","last_name":"Demol","created_date":"2018-06-11T21:39:49Z","updated_date":"2018-06-11T21:39:49Z","username":"email@email.com","email":"email@email.com","validated_email":false,"date_of_birth":"2018-06-04T00:06:00Z","account_enabled":false,"external_id":"d5b9b2f9-5e70-4c7a-b753-5833d3cc9253","custom_fields":[{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Registration type","FR":"Type d'inscription"},"values":{"EN":["new registration","renew"],"FR":["inscription","renouvellement"]},"position":0,"field_type":"INPUT_SELECT","id_str":"null"},"data":{"value":"inscription"}},{"field":{"enabled":true,"access_control":"PUBLIC","labels":{"EN":"Discipline","FR":"Discipline"},"placeholders":{"EN":"Discipline","FR":"Discipline"},"values":{"EN":["Natation","Plongeon","Triathlon","Waterpolo","Dauphin","Enf","Natation artistique"],"FR":["Natation","Plongeon","Triathlon","Waterpolo","Dauphin","Enf","Natation artistique"]},"position":1,"field_type":"INPUT_SELECT","id_str":"null"},"data":{"value":"Natation"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Nation","FR":"Nation"},"position":2,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":["ab", "cd"]}},{"field":{"enabled":true,"access_control":"PUBLIC","labels":{"EN":"Gender","FR":"Sexe"},"values":{"EN":["Male","Female"],"FR":["Homme","Femme"]},"position":3,"field_type":"INPUT_SELECT","id_str":"null"},"data":{"value":"Homme"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Mother firstname","FR":"Mère : prénom"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":4,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Brit"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Mother lastname","FR":"Mère : nom"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":5,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Dem"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Mother adress","FR":"Mère : adresse"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":6,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"123 rue de la Pompe"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Mother postal code - city","FR":"Mère : CP - Ville"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":7,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Haz"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Mother mobile number","FR":"Mère : mobile"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":8,"field_type":"INPUT_PHONE","id_str":"null"},"data":{"value":645454545}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Mother job title","FR":"Profession de la mère"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"placeholders":{"EN":"Mother job title","FR":"Profession de la mère"},"position":9,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Cut"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Mother company","FR":"Mère : profession"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":10,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Tot"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Father firstname","FR":"Père : prénom"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":11,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Yv"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Father lastname","FR":"Père : nom"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":12,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Dem"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Father address","FR":"Père : adresse"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":13,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"123 rue de la Pompe"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Father postal code - city","FR":"Père : code postal - ville"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":14,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Haz"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Father job title","FR":"Père : profession"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":15,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Cut"}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Father mobile","FR":"Père : mobile"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":16,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":646464646}},{"field":{"enabled":true,"access_control":"PRIVATE","labels":{"EN":"Father company","FR":"Père : employeur"},"descriptions":{"EN":"Pour mineur uniquement","FR":"Pour mineur uniquement"},"position":17,"id_str":"null","field_type":"INPUT_TEXT"},"data":{"value":"Cut"}}],"full_name":"Rémi Demol","displayed_name":"Rémi Demol","entity_type":"USER","id_str":"1267264340256770741","type":"User","authorities":["ROLE_USER"]}
}
`
)
