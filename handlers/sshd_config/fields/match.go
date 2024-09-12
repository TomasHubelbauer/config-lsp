package fields

var MatchAllowedOptions = map[string]struct{}{
	"AcceptEnv":                       {},
	"AllowAgentForwarding":            {},
	"AllowGroups":                     {},
	"AllowStreamLocalForwarding":      {},
	"AllowTcpForwarding":              {},
	"AllowUsers":                      {},
	"AuthenticationMethods":           {},
	"AuthorizedKeysCommand":           {},
	"AuthorizedKeysCommandUser":       {},
	"AuthorizedKeysFile":              {},
	"AuthorizedPrincipalsCommand":     {},
	"AuthorizedPrincipalsCommandUser": {},
	"AuthorizedPrincipalsFile":        {},
	"Banner":                          {},
	"CASignatureAlgorithms":           {},
	"ChannelTimeout":                  {},
	"ChrootDirectory":                 {},
	"ClientAliveCountMax":             {},
	"ClientAliveInterval":             {},
	"DenyGroups":                      {},
	"DenyUsers":                       {},
	"DisableForwarding":               {},
	"ExposeAuthInfo":                  {},
	"ForceCommand":                    {},
	"GatewayPorts":                    {},
	"GSSAPIAuthentication":            {},
	"HostbasedAcceptedAlgorithms":     {},
	"HostbasedAuthentication":         {},
	"HostbasedUsesNameFromPacketOnly": {},
	"IgnoreRhosts":                    {},
	"Include":                         {},
	"IPQoS":                           {},
	"KbdInteractiveAuthentication":    {},
	"KerberosAuthentication":          {},
	"LogLevel":                        {},
	"MaxAuthTries":                    {},
	"MaxSessions":                     {},
	"PasswordAuthentication":          {},
	"PermitEmptyPasswords":            {},
	"PermitListen":                    {},
	"PermitOpen":                      {},
	"PermitRootLogin":                 {},
	"PermitTTY":                       {},
	"PermitTunnel":                    {},
	"PermitUserRC":                    {},
	"PubkeyAcceptedAlgorithms":        {},
	"PubkeyAuthentication":            {},
	"PubkeyAuthOptions":               {},
	"RekeyLimit":                      {},
	"RevokedKeys":                     {},
	"RDomain":                         {},
	"SetEnv":                          {},
	"StreamLocalBindMask":             {},
	"StreamLocalBindUnlink":           {},
	"TrustedUserCAKeys":               {},
	"UnusedConnectionTimeout":         {},
	"X11DisplayOffset":                {},
	"X11Forwarding":                   {},
	"X11UseLocalhos":                  {},
}
