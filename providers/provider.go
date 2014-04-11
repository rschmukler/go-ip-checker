package providers

var providers map[string]Provider = make(map[string]Provider)

type Provider interface {
  Update(config map[string]string)
}

func GetProvider(name string) (Provider, bool) {
  result, present := providers[name]
  return result, present
}

func RegisterProvider(name string, provider Provider) {
  providers[name] = provider;
}

func ListProviders() (string) {
  list := ""
  for name, _ := range providers {
    list += "\t" + name +"\n"
  }
  return list
}