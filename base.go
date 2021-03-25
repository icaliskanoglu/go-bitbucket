package bitbucket

type Base struct {
	c *Client
}

func (u *Base) Get(t string) (interface{}, error) {

	urlStr := u.c.GetApiBaseURL() + "/" + t + "/"
	return u.c.execute("GET", urlStr, "")
}
