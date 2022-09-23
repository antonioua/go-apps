package config

import (
	"flag"
)

type Config struct {
	// Filed starts from uppercase letter allow accessing it
	FilePath       string
	TitleType      string
	PrimaryTitle   string
	OriginalTitle  string
	Genre          string
	EndYear        string
	RuntimeMinutes string
	Genres         string
	StartYear      string
	MaxRunTime     string
	MaxApiRequests int
	MaxRequests    int
	PlotFilter     string
}

func ParseConfig(args []string) (cf Config) {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.StringVar(&cf.FilePath, "file-path", "", "path to the inflated `title.basics.tsv.gz` file")
	fs.StringVar(&cf.TitleType, "title-type", "", "filter on `titleType` column")
	fs.StringVar(&cf.PrimaryTitle, "primary-title", "", "filter on `primaryTitle` column")
	fs.StringVar(&cf.OriginalTitle, "original-title", "", "filter on `originalTitle` column")
	fs.StringVar(&cf.Genre, "genre", "", "filter on `genre` column")
	fs.StringVar(&cf.StartYear, "start-year", "", "filter on `startYear` column")
	fs.StringVar(&cf.EndYear, "end-year", "", "filter on `endYear` column")
	fs.StringVar(&cf.RuntimeMinutes, "runtime-minutes", "", "filter on `runtimeMinutes` column")
	fs.StringVar(&cf.Genres, "genres", "", "filter on `genres` column")
	fs.IntVar(&cf.MaxApiRequests, "max-apiRequests", 1, "maximum number of requests to be made to [omdbapi](https://www.omdbapi.com/)")
	// Add validation for maxRunTime. Format should be https://pkg.go.dev/time#ParseDuration
	fs.StringVar(&cf.MaxRunTime, "max-runTime", "10m0s", "maximum run time of the application. Format is a `time.Duration` string, see (https://godoc.org/time#ParseDuration)")
	fs.IntVar(&cf.MaxRequests, "max-requests", 1, "maximum number of requests to send to https://www.omdbapi.com/")
	fs.StringVar(&cf.PlotFilter, "plot-filter", "", "regex pattern to apply to the plot of a film retrieved from https://www.omdbapi.com/")
	err := fs.Parse(args)
	if err != nil {
		return Config{}
	}
	return
}
