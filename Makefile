package:
	packer build -force -machine-readable -var "git_sha=`git rev-parse --short HEAD`"