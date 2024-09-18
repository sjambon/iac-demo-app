# Inleiding tot infrastructure as code met Terraform

## infrastructure/

De infrastrucure folder bevat alle terraform code waarmee we gaan werken.

Zie de [readme](./infrastructure/README.md).

## src/

Deze folder bevat de sourcecode van de applicaties die we in deze workshop zullen deployen.

Om de applicaties te bouwen:
- ```cd src/```
- ```make build```

Dit creert een folder ```src/output/``` met daarin 2 Azure function apps.