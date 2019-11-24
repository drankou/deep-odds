package types

import (
	"encoding/json"
	"log"
	"testing"
)

func TestRemoveDuplicitResultOdds(t *testing.T) {
	jsonResultOdds := `[
    {
        "id" : "43640609",
        "home_od" : "51.000",
        "draw_od" : "51.000",
        "away_od" : "1.002",
        "ss" : "0-2",
        "time_str" : "89",
        "add_time" : "1574531030"
    },
    {
        "id" : "43640502",
        "home_od" : "51.000",
        "draw_od" : "41.000",
        "away_od" : "1.004",
        "ss" : "0-2",
        "time_str" : "88",
        "add_time" : "1574530963"
    },
    {
        "id" : "43640442",
        "home_od" : "51.000",
        "draw_od" : "34.000",
        "away_od" : "1.005",
        "ss" : "0-2",
        "time_str" : "87",
        "add_time" : "1574530921"
    },
    {
        "id" : "43640408",
        "home_od" : "41.000",
        "draw_od" : "29.000",
        "away_od" : "1.007",
        "ss" : "0-2",
        "time_str" : "87",
        "add_time" : "1574530902"
    },
    {
        "id" : "43640343",
        "home_od" : "41.000",
        "draw_od" : "26.000",
        "away_od" : "1.008",
        "ss" : "0-2",
        "time_str" : "86",
        "add_time" : "1574530871"
    },
    {
        "id" : "43640340",
        "home_od" : "41.000",
        "draw_od" : "29.000",
        "away_od" : "1.007",
        "ss" : "0-2",
        "time_str" : "86",
        "add_time" : "1574530870"
    },
    {
        "id" : "43640182",
        "home_od" : "41.000",
        "draw_od" : "26.000",
        "away_od" : "1.008",
        "ss" : "0-2",
        "time_str" : "84",
        "add_time" : "1574530772"
    },
    {
        "id" : "43640075",
        "home_od" : "41.000",
        "draw_od" : "26.000",
        "away_od" : "1.010",
        "ss" : "0-2",
        "time_str" : "83",
        "add_time" : "1574530702"
    },
    {
        "id" : "43640014",
        "home_od" : "41.000",
        "draw_od" : "23.000",
        "away_od" : "1.012",
        "ss" : "0-2",
        "time_str" : "82",
        "add_time" : "1574530651"
    },
    {
        "id" : "43640007",
        "home_od" : "34.000",
        "draw_od" : "23.000",
        "away_od" : "1.012",
        "ss" : "0-2",
        "time_str" : "82",
        "add_time" : "1574530647"
    },
    {
        "id" : "43639996",
        "home_od" : "41.000",
        "draw_od" : "23.000",
        "away_od" : "1.012",
        "ss" : "0-2",
        "time_str" : "82",
        "add_time" : "1574530643"
    },
    {
        "id" : "43639955",
        "home_od" : "34.000",
        "draw_od" : "23.000",
        "away_od" : "1.012",
        "ss" : "0-2",
        "time_str" : "82",
        "add_time" : "1574530620"
    },
    {
        "id" : "43639827",
        "home_od" : "34.000",
        "draw_od" : "21.000",
        "away_od" : "1.015",
        "ss" : "0-2",
        "time_str" : "81",
        "add_time" : "1574530538"
    },
    {
        "id" : "43639662",
        "home_od" : "34.000",
        "draw_od" : "19.000",
        "away_od" : "1.020",
        "ss" : "0-2",
        "time_str" : "79",
        "add_time" : "1574530429"
    },
    {
        "id" : "43639626",
        "home_od" : "34.000",
        "draw_od" : "17.000",
        "away_od" : "1.025",
        "ss" : "0-2",
        "time_str" : "78",
        "add_time" : "1574530407"
    },
    {
        "id" : "43639427",
        "home_od" : "29.000",
        "draw_od" : "17.000",
        "away_od" : "1.025",
        "ss" : "0-2",
        "time_str" : "76",
        "add_time" : "1574530282"
    },
    {
        "id" : "43639291",
        "home_od" : "29.000",
        "draw_od" : "15.000",
        "away_od" : "1.030",
        "ss" : "0-2",
        "time_str" : "75",
        "add_time" : "1574530179"
    },
    {
        "id" : "43639233",
        "home_od" : "29.000",
        "draw_od" : "13.000",
        "away_od" : "1.036",
        "ss" : "0-2",
        "time_str" : "74",
        "add_time" : "1574530140"
    },
    {
        "id" : "43639188",
        "home_od" : "26.000",
        "draw_od" : "13.000",
        "away_od" : "1.036",
        "ss" : "0-2",
        "time_str" : "73",
        "add_time" : "1574530106"
    },
    {
        "id" : "43638968",
        "home_od" : "26.000",
        "draw_od" : "13.000",
        "away_od" : "1.040",
        "ss" : "0-2",
        "time_str" : "71",
        "add_time" : "1574529963"
    },
    {
        "id" : "43638889",
        "home_od" : "26.000",
        "draw_od" : "11.000",
        "away_od" : "1.045",
        "ss" : "0-2",
        "time_str" : "70",
        "add_time" : "1574529901"
    },
    {
        "id" : "43638885",
        "home_od" : "23.000",
        "draw_od" : "11.000",
        "away_od" : "1.050",
        "ss" : "0-2",
        "time_str" : "70",
        "add_time" : "1574529898"
    },
    {
        "id" : "43638880",
        "home_od" : "26.000",
        "draw_od" : "11.000",
        "away_od" : "1.045",
        "ss" : "0-2",
        "time_str" : "70",
        "add_time" : "1574529896"
    },
    {
        "id" : "43638669",
        "home_od" : "23.000",
        "draw_od" : "11.000",
        "away_od" : "1.050",
        "ss" : "0-2",
        "time_str" : "67",
        "add_time" : "1574529735"
    },
    {
        "id" : "43638609",
        "home_od" : "21.000",
        "draw_od" : "11.000",
        "away_od" : "1.050",
        "ss" : "0-2",
        "time_str" : "67",
        "add_time" : "1574529698"
    },
    {
        "id" : "43638537",
        "home_od" : "21.000",
        "draw_od" : "10.000",
        "away_od" : "1.056",
        "ss" : "0-2",
        "time_str" : "66",
        "add_time" : "1574529640"
    },
    {
        "id" : "43638375",
        "home_od" : "21.000",
        "draw_od" : "10.000",
        "away_od" : "1.062",
        "ss" : "0-2",
        "time_str" : "64",
        "add_time" : "1574529514"
    },
    {
        "id" : "43638222",
        "home_od" : "19.000",
        "draw_od" : "10.000",
        "away_od" : "1.062",
        "ss" : "0-2",
        "time_str" : "62",
        "add_time" : "1574529403"
    },
    {
        "id" : "43638167",
        "home_od" : "19.000",
        "draw_od" : "9.500",
        "away_od" : "1.071",
        "ss" : "0-2",
        "time_str" : "61",
        "add_time" : "1574529366"
    },
    {
        "id" : "43638136",
        "home_od" : "17.000",
        "draw_od" : "9.500",
        "away_od" : "1.071",
        "ss" : "0-2",
        "time_str" : "61",
        "add_time" : "1574529347"
    },
    {
        "id" : "43637925",
        "home_od" : "17.000",
        "draw_od" : "9.000",
        "away_od" : "1.071",
        "ss" : "0-2",
        "time_str" : "59",
        "add_time" : "1574529223"
    },
    {
        "id" : "43637825",
        "home_od" : "17.000",
        "draw_od" : "9.000",
        "away_od" : "1.083",
        "ss" : "0-2",
        "time_str" : "58",
        "add_time" : "1574529163"
    },
    {
        "id" : "43637770",
        "home_od" : "15.000",
        "draw_od" : "9.000",
        "away_od" : "1.083",
        "ss" : "0-2",
        "time_str" : "57",
        "add_time" : "1574529130"
    },
    {
        "id" : "43637645",
        "home_od" : "15.000",
        "draw_od" : "8.500",
        "away_od" : "1.083",
        "ss" : "0-2",
        "time_str" : "56",
        "add_time" : "1574529044"
    },
    {
        "id" : "43637410",
        "home_od" : "15.000",
        "draw_od" : "8.500",
        "away_od" : "1.091",
        "ss" : "0-2",
        "time_str" : "53",
        "add_time" : "1574528902"
    },
    {
        "id" : "43637346",
        "home_od" : "13.000",
        "draw_od" : "8.500",
        "away_od" : "1.100",
        "ss" : "0-2",
        "time_str" : "53",
        "add_time" : "1574528859"
    },
    {
        "id" : "43637113",
        "home_od" : "13.000",
        "draw_od" : "8.000",
        "away_od" : "1.100",
        "ss" : "0-2",
        "time_str" : "50",
        "add_time" : "1574528725"
    },
    {
        "id" : "43637084",
        "home_od" : "13.000",
        "draw_od" : "8.000",
        "away_od" : "1.111",
        "ss" : "0-2",
        "time_str" : "50",
        "add_time" : "1574528700"
    },
    {
        "id" : "43636805",
        "home_od" : "12.000",
        "draw_od" : "8.000",
        "away_od" : "1.111",
        "ss" : "0-2",
        "time_str" : "47",
        "add_time" : "1574528504"
    },
    {
        "id" : "43636781",
        "home_od" : "12.000",
        "draw_od" : "8.000",
        "away_od" : "1.125",
        "ss" : "0-1",
        "time_str" : "46",
        "add_time" : "1574528487"
    },
    {
        "id" : "43636751",
        "home_od" : "-",
        "draw_od" : "-",
        "away_od" : "-",
        "ss" : "0-1",
        "time_str" : "46",
        "add_time" : "1574528468"
    },
    {
        "id" : "43636737",
        "home_od" : "6.000",
        "draw_od" : "4.000",
        "away_od" : "1.500",
        "ss" : "0-1",
        "time_str" : "46",
        "add_time" : "1574528458"
    },
    {
        "id" : "43634759",
        "home_od" : "5.500",
        "draw_od" : "4.000",
        "away_od" : "1.500",
        "ss" : "0-1",
        "time_str" : "45",
        "add_time" : "1574527632"
    },
    {
        "id" : "43634414",
        "home_od" : "7.000",
        "draw_od" : "4.333",
        "away_od" : "1.400",
        "ss" : "0-1",
        "time_str" : "45",
        "add_time" : "1574527552"
    },
    {
        "id" : "43633959",
        "home_od" : "6.500",
        "draw_od" : "4.333",
        "away_od" : "1.444",
        "ss" : "0-0",
        "time_str" : "44",
        "add_time" : "1574527453"
    },
    {
        "id" : "43633814",
        "home_od" : "-",
        "draw_od" : "-",
        "away_od" : "-",
        "ss" : "0-0",
        "time_str" : "43",
        "add_time" : "1574527420"
    },
    {
        "id" : "43633401",
        "home_od" : "2.625",
        "draw_od" : "3.000",
        "away_od" : "2.625",
        "ss" : "0-0",
        "time_str" : "42",
        "add_time" : "1574527312"
    },
    {
        "id" : "43633219",
        "home_od" : "2.625",
        "draw_od" : "3.000",
        "away_od" : "2.600",
        "ss" : "0-0",
        "time_str" : "41",
        "add_time" : "1574527265"
    },
    {
        "id" : "43632267",
        "home_od" : "2.625",
        "draw_od" : "3.100",
        "away_od" : "2.600",
        "ss" : "0-0",
        "time_str" : "37",
        "add_time" : "1574527028"
    },
    {
        "id" : "43632218",
        "home_od" : "2.600",
        "draw_od" : "3.100",
        "away_od" : "2.600",
        "ss" : "0-0",
        "time_str" : "36",
        "add_time" : "1574527007"
    },
    {
        "id" : "43631953",
        "home_od" : "2.600",
        "draw_od" : "3.200",
        "away_od" : "2.600",
        "ss" : "0-0",
        "time_str" : "35",
        "add_time" : "1574526941"
    },
    {
        "id" : "43631748",
        "home_od" : "2.600",
        "draw_od" : "3.200",
        "away_od" : "2.500",
        "ss" : "0-0",
        "time_str" : "34",
        "add_time" : "1574526879"
    },
    {
        "id" : "43630936",
        "home_od" : "2.600",
        "draw_od" : "3.250",
        "away_od" : "2.500",
        "ss" : "0-0",
        "time_str" : "30",
        "add_time" : "1574526644"
    },
    {
        "id" : "43630904",
        "home_od" : "2.600",
        "draw_od" : "3.400",
        "away_od" : "2.500",
        "ss" : "0-0",
        "time_str" : "30",
        "add_time" : "1574526634"
    },
    {
        "id" : "43630800",
        "home_od" : "2.600",
        "draw_od" : "3.250",
        "away_od" : "2.500",
        "ss" : "0-0",
        "time_str" : "30",
        "add_time" : "1574526601"
    },
    {
        "id" : "43629831",
        "home_od" : "2.600",
        "draw_od" : "3.400",
        "away_od" : "2.500",
        "ss" : "0-0",
        "time_str" : "24",
        "add_time" : "1574526265"
    },
    {
        "id" : "43629761",
        "home_od" : "2.600",
        "draw_od" : "3.500",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "24",
        "add_time" : "1574526243"
    },
    {
        "id" : "43629705",
        "home_od" : "2.600",
        "draw_od" : "3.400",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "23",
        "add_time" : "1574526224"
    },
    {
        "id" : "43629640",
        "home_od" : "2.600",
        "draw_od" : "3.500",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "23",
        "add_time" : "1574526203"
    },
    {
        "id" : "43629629",
        "home_od" : "2.500",
        "draw_od" : "3.500",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "23",
        "add_time" : "1574526200"
    },
    {
        "id" : "43629570",
        "home_od" : "2.500",
        "draw_od" : "3.400",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "23",
        "add_time" : "1574526175"
    },
    {
        "id" : "43629105",
        "home_od" : "2.600",
        "draw_od" : "3.500",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "20",
        "add_time" : "1574526008"
    },
    {
        "id" : "43628595",
        "home_od" : "2.500",
        "draw_od" : "3.500",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "17",
        "add_time" : "1574525822"
    },
    {
        "id" : "43628416",
        "home_od" : "2.600",
        "draw_od" : "3.500",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "16",
        "add_time" : "1574525760"
    },
    {
        "id" : "43628274",
        "home_od" : "2.600",
        "draw_od" : "3.500",
        "away_od" : "2.375",
        "ss" : "0-0",
        "time_str" : "15",
        "add_time" : "1574525714"
    },
    {
        "id" : "43628198",
        "home_od" : "2.600",
        "draw_od" : "3.600",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "14",
        "add_time" : "1574525681"
    },
    {
        "id" : "43628175",
        "home_od" : "2.600",
        "draw_od" : "3.500",
        "away_od" : "2.375",
        "ss" : "0-0",
        "time_str" : "14",
        "add_time" : "1574525672"
    },
    {
        "id" : "43628148",
        "home_od" : "2.600",
        "draw_od" : "3.600",
        "away_od" : "2.375",
        "ss" : "0-0",
        "time_str" : "14",
        "add_time" : "1574525662"
    },
    {
        "id" : "43628039",
        "home_od" : "2.600",
        "draw_od" : "3.500",
        "away_od" : "2.375",
        "ss" : "0-0",
        "time_str" : "13",
        "add_time" : "1574525624"
    },
    {
        "id" : "43627743",
        "home_od" : "2.600",
        "draw_od" : "3.600",
        "away_od" : "2.375",
        "ss" : "0-0",
        "time_str" : "12",
        "add_time" : "1574525526"
    },
    {
        "id" : "43627677",
        "home_od" : "2.600",
        "draw_od" : "3.600",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "12",
        "add_time" : "1574525510"
    },
    {
        "id" : "43627652",
        "home_od" : "2.600",
        "draw_od" : "3.600",
        "away_od" : "2.375",
        "ss" : "0-0",
        "time_str" : "11",
        "add_time" : "1574525502"
    },
    {
        "id" : "43626610",
        "home_od" : "2.500",
        "draw_od" : "3.600",
        "away_od" : "2.400",
        "ss" : "0-0",
        "time_str" : "4",
        "add_time" : "1574525065"
    },
    {
        "id" : "43626356",
        "home_od" : "2.500",
        "draw_od" : "3.600",
        "away_od" : "2.375",
        "ss" : "0-0",
        "time_str" : "2",
        "add_time" : "1574524918"
    },
    {
        "id" : "43625558",
        "home_od" : "2.500",
        "draw_od" : "3.600",
        "away_od" : "2.300",
        "ss" : "",
        "time_str" : "",
        "add_time" : "1574524122"
    },
    {
        "id" : "43536171",
        "home_od" : "2.400",
        "draw_od" : "3.600",
        "away_od" : "2.400",
        "ss" : "",
        "time_str" : "",
        "add_time" : "1574421929"
    }
]`

	var resultOdds []Result
	err := json.Unmarshal([]byte(jsonResultOdds), &resultOdds)
	if err != nil {
		t.Fatal(err)
	}

	log.Print("Result odds length: ", len(resultOdds))

	res := RemoveDuplicitResultOdds(resultOdds)
	log.Print("Clean odds length: ", len(res))
	log.Printf("%+v", res)
}
