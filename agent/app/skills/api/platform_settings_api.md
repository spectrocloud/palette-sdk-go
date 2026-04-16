# platform_settings_api

Configure platform settings: SSO, session timeout, login banner, password policy, FIPS

**get_saml** — Get SAML SSO configuration
`{"mode": "get_saml", "tenant_uid": "tenant-uid-here"}`

**update_saml** — Update SAML SSO configuration
`{"mode": "update_saml", "tenant_uid": "tenant-uid", "spec": {"idpMetadataUrl": "https://idp.example.com/metadata"}}`

**get_session_timeout** — Get session timeout settings
`{"mode": "get_session_timeout", "tenant_uid": "tenant-uid-here"}`

**update_session_timeout** — Update session timeout settings
`{"mode": "update_session_timeout", "tenant_uid": "tenant-uid", "spec": {"accessTokenExpiry": 60, "refreshTokenExpiry": 1440}}`

**get_login_banner** — Get login banner settings
`{"mode": "get_login_banner", "tenant_uid": "tenant-uid-here"}`

**update_login_banner** — Update login banner settings
`{"mode": "update_login_banner", "tenant_uid": "tenant-uid", "spec": {"message": "Authorized users only", "isEnabled": true}}`

**get_oidc** — Get OIDC SSO configuration
`{"mode": "get_oidc", "tenant_uid": "tenant-uid-here"}`

**get_password_policy** — Get password policy
`{"mode": "get_password_policy", "tenant_uid": "tenant-uid-here"}`

**update_password_policy** — Update password policy
`{"mode": "update_password_policy", "tenant_uid": "tenant-uid", "spec": {"minLength": 12, "requireUppercase": true, "requireNumber": true}}`

**get_fips** — Get FIPS compliance preference
`{"mode": "get_fips", "tenant_uid": "tenant-uid-here"}`
*FIPS cannot be disabled once enabled*

**update_fips** — Enable FIPS compliance (cannot be undone)
`{"mode": "update_fips", "tenant_uid": "tenant-uid", "spec": {"isFipsEnabled": true}}`
*WARNING: FIPS cannot be disabled once enabled*

**get_tenant_uid** — Get the current tenant UID
`{"mode": "get_tenant_uid"}`

**update_oidc** — Update OIDC SSO configuration
`{"mode": "update_oidc", "tenant_uid": "tenant-uid", "spec": {"issuerUrl": "https://idp.example.com", "clientId": "...", "clientSecret": "..."}}`

