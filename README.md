go build -o terraform-provider-envtrack .

mkdir -p ~/.terraform.d/plugins/local.providers/local/envtrack/1.0.4/darwin_arm64/

cp terraform-provider-envtrack ~/.terraform.d/plugins/local.providers/local/envtrack/1.0.3/darwin_arm64/
