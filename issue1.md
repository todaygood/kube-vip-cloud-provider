[root@pcentos kube-vip-cloud-provider]# docker build . -t pastack-registry.paic.com.cn/kube-vip-cloud-provider:0.0.7.1
[+] Building 197.1s (13/14)
 => [internal] load .dockerignore                                                                                                            0.0s
  => => transferring context: 2B                                                                                                              0.0s
   => [internal] load build definition from Dockerfile                                                                                         0.0s
    => => transferring dockerfile: 939B                                                                                                         0.0s
	 => resolve image config for docker.io/docker/dockerfile:1.4                                                                                21.2s
	  => CACHED docker-image://docker.io/docker/dockerfile:1.4@sha256:9ba7531bd80fb0a858632727cf7a112fbfd19b17e94c4e84ced81e24ef1a0dbc            0.0s
	   => [internal] load metadata for docker.io/library/golang:1.19                                                                              47.5s
	    => [internal] load metadata for gcr.m.daocloud.io/distroless/static:nonroot                                                                 2.5s
		 => [internal] load build context                                                                                                            0.0s
		  => => transferring context: 8.34kB                                                                                                          0.0s
		   => CACHED [stage-1 1/2] FROM gcr.m.daocloud.io/distroless/static:nonroot@sha256:9ecc53c269509f63c69a266168e4a687c7eb8c0cfd753bd8bfcaa4f58a  0.0s
		    => CACHED [builder 1/5] FROM docker.io/library/golang:1.19@sha256:d680597c753e7c73814ddd09c69bef5b78922d81d3ca103c82fa69771ea8151c          0.0s
			 => [builder 2/5] COPY . /src/                                                                                                               0.1s
			  => [builder 3/5] WORKDIR /src                                                                                                               0.1s
			   => [builder 4/5] RUN  --mount=type=cache,target=/root/.local/share/golang     --mount=type=cache,target=/go/pkg/mod     go mod download   121.3s
			    => ERROR [builder 5/5] RUN --mount=type=cache,target=/root/.cache/go-build     --mount=type=cache,target=/go/pkg/mod     --mount=type=cach  6.3s
				------
				 > [builder 5/5] RUN --mount=type=cache,target=/root/.cache/go-build     --mount=type=cache,target=/go/pkg/mod     --mount=type=cache,target=/root/.local/share/golang     CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w -extldflags -static" -o kube-vip-cloud-provider:
				 #8 6.238 /go/pkg/mod/k8s.io/apiserver@v0.25.4/pkg/storage/value/encrypt/envelope/envelope.go:33:2: missing go.sum entry for module providing package golang.org/x/crypto/cryptobyte (imported by k8s.io/apiserver/pkg/storage/value/encrypt/envelope); to add:
				 #8 6.238        go get k8s.io/apiserver/pkg/storage/value/encrypt/envelope@v0.25.4
				 #8 6.239 /go/pkg/mod/k8s.io/apiserver@v0.25.4/pkg/storage/value/encrypt/secretbox/secretbox.go:25:2: missing go.sum entry for module providing package golang.org/x/crypto/nacl/secretbox (imported by k8s.io/apiserver/pkg/storage/value/encrypt/secretbox); to add:
				 #8 6.239        go get k8s.io/apiserver/pkg/storage/value/encrypt/secretbox@v0.25.4
				 #8 6.239 /go/pkg/mod/github.com/grpc-ecosystem/go-grpc-prometheus@v1.2.0/client_metrics.go:7:2: missing go.sum entry for module providing package golang.org/x/net/context (imported by github.com/grpc-ecosystem/go-grpc-prometheus); to add:
				 #8 6.239        go get github.com/grpc-ecosystem/go-grpc-prometheus@v1.2.0
				 #8 6.239 /go/pkg/mod/golang.org/x/oauth2@v0.0.0-20211104180415-d3ed0bb246c8/internal/token.go:23:2: missing go.sum entry for module providing package golang.org/x/net/context/ctxhttp (imported by golang.org/x/oauth2/internal); to add:
				 #8 6.239        go get golang.org/x/oauth2/internal@v0.0.0-20211104180415-d3ed0bb246c8
				 #8 6.239 /go/pkg/mod/k8s.io/apimachinery@v0.25.4/pkg/util/net/http.go:39:2: missing go.sum entry for module providing package golang.org/x/net/http2 (imported by k8s.io/client-go/rest); to add:
				 #8 6.239        go get k8s.io/client-go/rest@v0.25.4
				 #8 6.239 /go/pkg/mod/google.golang.org/grpc@v1.47.0/internal/transport/controlbuf.go:31:2: missing go.sum entry for module providing package golang.org/x/net/http2/hpack (imported by google.golang.org/grpc/internal/transport); to add:
				 #8 6.239        go get google.golang.org/grpc/internal/transport@v1.47.0
				 #8 6.239 /go/pkg/mod/github.com/!puerkito!bio/purell@v1.1.1/purell.go:17:2: missing go.sum entry for module providing package golang.org/x/net/idna (imported by github.com/PuerkitoBio/purell); to add:
				 #8 6.239        go get github.com/PuerkitoBio/purell@v1.1.1
				 #8 6.239 /go/pkg/mod/google.golang.org/grpc@v1.47.0/server.go:36:2: missing go.sum entry for module providing package golang.org/x/net/trace (imported by google.golang.org/grpc); to add:
				 #8 6.239        go get google.golang.org/grpc@v1.47.0
				 #8 6.239 /go/pkg/mod/k8s.io/apiserver@v0.25.4/pkg/util/wsstream/conn.go:28:2: missing go.sum entry for module providing package golang.org/x/net/websocket (imported by k8s.io/apiserver/pkg/endpoints/handlers); to add:
				 #8 6.239        go get k8s.io/apiserver/pkg/endpoints/handlers@v0.25.4
				 #8 6.239 /go/pkg/mod/k8s.io/apiserver@v0.25.4/pkg/authentication/token/cache/cached_token_authenticator.go:33:2: missing go.sum entry for module providing package golang.org/x/sync/singleflight (imported by k8s.io/apiserver/pkg/authentication/token/cache); to add:
				 #8 6.239        go get k8s.io/apiserver/pkg/authentication/token/cache@v0.25.4
				 #8 6.239 /go/pkg/mod/github.com/prometheus/procfs@v0.7.3/proc_maps.go:25:2: missing go.sum entry for module providing package golang.org/x/sys/unix (imported by k8s.io/apiserver/pkg/server/options); to add:
				 #8 6.239        go get k8s.io/apiserver/pkg/server/options@v0.25.4
				 #8 6.239 /go/pkg/mod/k8s.io/client-go@v0.25.4/plugin/pkg/client/auth/exec/exec.go:36:2: missing go.sum entry for module providing package golang.org/x/term (imported by k8s.io/client-go/tools/clientcmd); to add:
				 #8 6.239        go get k8s.io/client-go/tools/clientcmd@v0.25.4
				 #8 6.239 /go/pkg/mod/github.com/!puerkito!bio/purell@v1.1.1/purell.go:18:2: missing go.sum entry for module providing package golang.org/x/text/unicode/norm (imported by github.com/PuerkitoBio/purell); to add:
				 #8 6.239        go get github.com/PuerkitoBio/purell@v1.1.1
				 #8 6.239 /go/pkg/mod/github.com/!puerkito!bio/purell@v1.1.1/purell.go:19:2: missing go.sum entry for module providing package golang.org/x/text/width (imported by github.com/PuerkitoBio/purell); to add:
				 #8 6.239        go get github.com/PuerkitoBio/purell@v1.1.1
				 ------
				 process "/bin/sh -c CGO_ENABLED=0 GOOS=linux go build -ldflags \"$LD_FLAGS\" -o kube-vip-cloud-provider" did not complete successfully: exit code: 1
