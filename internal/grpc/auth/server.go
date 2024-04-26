package auth

import ssov1 "github.com/BakeNecko/GO_PROTOS/protos/gen/go/sso"

type serverAPI interface {
	ssov1.UnimplementedAuthServer
}

func RegisterServerAPI(s *ssov1.Seerver, api *serverAPI) {
	s.Regiser
}
