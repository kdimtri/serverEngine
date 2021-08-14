package engine

import "context"

// MuontPoint type holds prefixes
// for private and public paths on cdn
type MuontPoint struct {
	private string
	public  string
}

var mnt *MuontPoint = &MuontPoint{}

// NewMountPoint sets persistent(for server lifetime or until new call for set method)
// global MuontPoint, and returns pointer to instance
func NewMountPoint(privatePath, publicPath string) *MuontPoint {
	mnt.private = privatePath
	mnt.public = publicPath
	return mnt
}

// SetMountPointFromContex same as NewMountPoint,
// only setter paths values being parsed from context
func SetMountPointFromContex(ctx context.Context) bool {
	public, ok := ctx.Value(ContextPublicMountPath).(string)
	if !ok {
		return ok
	}
	private, ok := ctx.Value(ContextPrivateMountPath).(string)
	if !ok {
		return ok
	}
	mnt.private = private
	mnt.public = public
	return true
}

// SetContext stores values of mountPoint to context
func (m *MuontPoint) SetContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, ContextPublicMountPath, m.public)
	ctx = context.WithValue(ctx, ContextPrivateMountPath, m.private)
	return ctx
}

// GetMountPointsFromContex func to get values of mountPoint from context
func GetMountPointsFromContex(ctx context.Context) (private string, public string) {
	public, ok := ctx.Value(ContextPublicMountPath).(string)
	if !ok {
		return
	}
	private, ok = ctx.Value(ContextPrivateMountPath).(string)
	if !ok {
		return
	}
	return
}

/*
func SetMountPoint(privatePath, publicPath string) bool {
	mnt.private = privatePath
	mnt.public = publicPath
	return true
}


func GetPrivateMount() string {
	return mnt.private
}
func GetPublicMount() string {
	return mnt.public
}

func (m *MuontPoint) SetRequestContext(r *http.Request) *http.Request {
	ctx := r.Context()
	ctx = m.SetContext(ctx)
	return r.WithContext(ctx)
}
func CutPrefix(s string) string {
	if strings.Contains(s, mnt.private) {
		return strings.TrimSuffix(s, mnt.private)
	}
	if strings.Contains(s, mnt.public) {
		return strings.TrimSuffix(s, mnt.public)
	}
	return s
}
func (m *MuontPoint) PrivateHTTPPath(path string) http.Dir {
	return http.Dir(m.private + path)
}
func (m *MuontPoint) PublicHTTPPath(path string) http.Dir {
	return http.Dir(m.public + path)
}
func (m *MuontPoint) PrivatePrefix(path string) string {
	return fmt.Sprintf("%v", m.PrivateHTTPPath(path))
}
func (m *MuontPoint) PublicPrefix(path string) string {
	return fmt.Sprintf("%v", m.PublicHTTPPath(path))
}
*/
