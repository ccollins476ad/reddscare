package main

type Config struct {
	BDFRDir string `yaml:"bdfr_dir"`
	WorkDir string `yaml:"work_dir"`
	HTMLDir string `yaml:"html_dir"`

	BDFROpts []string `yaml:"bdfr_opts"`

	BDFRscrapePath string   `yaml:"bdfrscrape_path"`
	BDFRscrapeOpts []string `yaml:"bdfrscrape_opts"`

	BDFRToHTMLOpts []string `yaml:"bdfrtohtml_opts"`
}

func (cfg *Config) IsComplete() bool {
	return cfg.BDFRDir != "" && cfg.WorkDir != "" && cfg.HTMLDir != "" && cfg.BDFRscrapePath != ""
}
