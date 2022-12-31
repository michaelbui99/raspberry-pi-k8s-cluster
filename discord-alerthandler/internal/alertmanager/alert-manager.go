package alertmanager

type AlertManagerDTO struct {
	Alerts            []AlertManagerAlert `json:"alerts"`
	CommonAnnotations struct {
		Summary string `json:"summary"`
	} `json:"commonAnnotations"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
	} `json:"commonLabels"`
	ExternalURL string `json:"externalURL"`
	GroupKey    string `json:"groupKey"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Version  string `json:"version"`
}

type AlertManagerAlert struct {
	Annotations  AlertManagerAnnotations `json:"annotations"`
	EndsAt       string                  `json:"endsAt"`
	GeneratorURL string                  `json:"generatorURL"`
	Labels       map[string]string       `json:"labels"`
	StartsAt     string                  `json:"startsAt"`
	Status       string                  `json:"status"`
}

type AlertManagerAnnotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

func NewAlertManagerAlert(
	annotations AlertManagerAnnotations,
	endsAt string,
	generatorUrl string,
	labels map[string]string,
	startsArt string,
	status string) *AlertManagerAlert {

	return &AlertManagerAlert{
		Annotations:  annotations,
		EndsAt:       endsAt,
		GeneratorURL: generatorUrl,
		Labels:       labels,
		StartsAt:     startsArt,
		Status:       status,
	}
}
