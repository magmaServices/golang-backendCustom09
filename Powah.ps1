# First, grant ExecutionPolicy on running PS scripts first (run PS as admin)
# PS> Set-ExecutionPolicy RemoteSigned
# https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_scripts?view=powershell-5.1
# Feel free to close PS for administrator and open new PS session in fesl-backend directory
# PS> cd $env:GOPATH/src/localhost/go-heroes/fesl-backend
#
# Then, dot-source the file in PS
# PS> . .\Powah.ps1
#
# Finally, run functions in PS like separate programs:
# i.e. type "MakeRun" and hit enter
# PS> MakeRun

$name = "fesl-backend"
$binary = ".\fesl-backend.exe"
$entrypoint = ".\cmd\${name}"

function MakeDeps {
    [CmdletBinding()]
    param()

    Write-Output "Downloading dependencies via docker..."
	glide.exe install
}

function MakeCodegen {
    [CmdletBinding()]
    param()

    Write-Output "Generating code..."
	go.exe generate $entrypoint
}

function MakeSeed {
    [CmdletBinding()]
    param()

    Write-Output "Seeding data..."
	go.exe run .\cmd\create-hero\main.go
}

function MakeCompile {
    [CmdletBinding()]
    param()

	Write-Output "Compiling Golang code as binary..."
	$env:CGO_ENABLED = 0;
    go.exe build -a -ldflags='-w -s' -v -o $binary $entrypoint
}

function MakeRun {
    [CmdletBinding()]
    param()

	Write-Output "Starting compiled binary..."
    go build -v -o $binary $entrypoint; if ($?) {
        $binargs = "-config", "dev.env"
        & $binary $binargs
    }
}

function MakeClean {
    [CmdletBinding()]
    param()

	Write-Output "Removing compiled binary..."
    Remove-Item $binary
}
