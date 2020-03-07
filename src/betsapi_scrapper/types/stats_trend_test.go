package types

import (
	"betsapi_scrapper/utils"
	"github.com/wcharczuk/go-chart"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

func TestAddMissingStatsTrend(t *testing.T) {
	data, err := ioutil.ReadFile(path.Join(utils.GetAbsPathToRoot(), "mock_data", "stats_trend.bson"))
	var statsTrend StatsTrend
	err = bson.Unmarshal(data, &statsTrend)
	if err != nil {
		log.Fatal(err)
	}

	newStats := AddMissingStatsTrend(&statsTrend)
	attacksHome := newStats.Attacks.Home
	attacksAway := newStats.Attacks.Away

	var attacksMinutes []float64
	var attacksHomeValues []float64
	for _, tick := range attacksHome {
		attacksMinutes = append(attacksMinutes, float64(tick.Time))
		attacksHomeValues = append(attacksHomeValues, float64(tick.Value))
	}

	var attacksAwayValues []float64
	for _, tick := range attacksAway {
		attacksAwayValues = append(attacksAwayValues, float64(tick.Value))
	}


	log.Print("Home attacks", attacksHomeValues)
	log.Print("Away attacks", attacksAwayValues)
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: attacksMinutes[:90],
				YValues: attacksHomeValues[:90],
			},
			chart.ContinuousSeries{
				XValues: attacksMinutes[:90],
				YValues: attacksAwayValues[:90],
			},
		},
	}

	pngFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}

	if err := graph.Render(chart.PNG, pngFile); err != nil {
		panic(err)
	}

	if err := pngFile.Close(); err != nil {
		panic(err)
	}
}
