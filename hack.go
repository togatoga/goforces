package goforces

type Hack struct {
	CreationTimeSeconds int64         `json:"creationTimeSeconds"`
	Defender            Party         `json:"defender"`
	Hacker              Party         `json:"hacker"`
	ID                  int64         `json:"id"`
	JudgeProtocol       JudgeProtocol `json:"judgeProtocol"`
	Problem             Problem       `json:"problem"`
	Test                string        `json:"test"`
	Verdict             string        `json:"verdict"`
}

type JudgeProtocol struct {
	Manual   string `json:"manual"`
	Protocol string `json:"protocol"`
	Verdict  string `json:"verdict"`
}
