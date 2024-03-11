.PHONY: ent
ent: 
	make mkdirent && make genent

.PHONY: mkdirent
mkdirent:
	@mkdir -p ./gen/ent && echo "package generate" > ./gen/generate.go

.PHONY: genent
genent:
ifneq ("$(wildcard ./gen/ent)","")
	@go run -mod=mod entgo.io/ent/cmd/ent generate \
    --feature privacy \
    --feature sql/modifier \
    --feature entql \
    --feature sql/upsert \
    --feature intercept \
    --feature schema/snapshot \
    --target ./gen/ent \
    ./schema
endif


