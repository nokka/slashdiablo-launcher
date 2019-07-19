# Use run if there's only changes to the QML layer.
run:
	./deploy/darwin/slashdiablo-launcher.app/Contents/MacOS/slashdiablo-launcher

# Rebuild the app if you made changes to the Go layer.
build:
	qtdeploy build