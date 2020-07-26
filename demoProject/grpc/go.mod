module grpc

go 1.12

require (
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
	zonst/logging v0.0.0-00010101000000-000000000000
	zonst/qipai-oversea/apisrv/propscommonapisrv v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20181219182458-5a97ab628bfb
	zonst/eagle => ../../././zonst/eagle
	zonst/logging => ../../././zonst/logging
	zonst/pool => ../../././zonst/pool
	zonst/qipai-oversea/apisrv/propscommonapisrv => ../../././zonst/qipai-oversea/apisrv/propscommonapisrv
	zonst/qipai-oversea/apisrv/utils => ../../././zonst/qipai-oversea/apisrv/utils
	zonst/qipai-oversea/lib/constants => ../../././zonst/qipai-oversea/lib/constants
	zonst/qipai-oversea/lib/contextutil => ../../././zonst/qipai-oversea/lib/contextutil
	zonst/qipai-oversea/lib/fsnotifyutil => ../../././zonst/qipai-oversea/lib/fsnotifyutil
	zonst/qipai-oversea/lib/gin-middlewares/secretauth => ../../././zonst/qipai-oversea/lib/gin-middlewares/secretauth
	zonst/qipai-oversea/lib/gin-middlewares/tokenauth => ../../././zonst/qipai-oversea/lib/gin-middlewares/tokenauth
	zonst/qipai-oversea/lib/logging => ../../././zonst/qipai-oversea/lib/logging
	zonst/qipai-oversea/lib/messages => ../../././zonst/qipai-oversea/lib/messages
	zonst/qipai-oversea/lib/poolutil => ../../././zonst/qipai-oversea/lib/poolutil
	zonst/qipai-oversea/lib/protocolutil => ../../././zonst/qipai-oversea/lib/protocolutil
	zonst/qipai-oversea/lib/structutil => ../../././zonst/qipai-oversea/lib/structutil
)
