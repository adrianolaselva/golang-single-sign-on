# Oauth2

Obs: client_id and client_secret passado como basic auth

Basic MTNkNjkzZDAtZDY2Yy0xMWU5LTg5YjctZGY1YjRkZjBhZmIzOjh6ZDlVTG1hNnhOTjF3YlI3aDhlcjd6N3FiRVJVTHNqQ3FEMnB6VDU=

## Flows

#### grant_type `password`

>https://curity.io/resources/develop/oauth/oauth-resource-owner-password-credential-flow/

```sh
curl -XPOST 'https://api.aglaia.io/oauth/token' \
    -H 'authorization: Basic MTNkYjA2ODAtZDY2Yy0xMWU5LTk4MDMtMzE0ZjYwOTFmZWVlOkltbGVVUUNCM0lORW9PNVVBU3AzY3FiM2ZBYkxVc2VJcm8wU0NxcGQ=' \
    -H 'content-type: application/x-www-form-urlencoded' \
    --data 'grant_type=password&scope=wm-api:admin+wm-api:integration&username=adrianolaselva@gmail.com&password=123@mudar' \
    --compressed
```

```json
{
   "token_type":"Bearer",
   "expires_in":1296000,
   "access_token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjIzNzA1Yjg3YTA2YThiZDJiYTkwNjRlZjg5NzY0NzBjMTJhODQ4Mjg3MzNjOWVjNDA1NzMwNjYzYWFkOTdhMzE5NWVkOWRkMDkyNmNhOTJjIn0.eyJhdWQiOiIxM2RiMDY4MC1kNjZjLTExZTktOTgwMy0zMTRmNjA5MWZlZWUiLCJqdGkiOiIyMzcwNWI4N2EwNmE4YmQyYmE5MDY0ZWY4OTc2NDcwYzEyYTg0ODI4NzMzYzllYzQwNTczMDY2M2FhZDk3YTMxOTVlZDlkZDA5MjZjYTkyYyIsImlhdCI6MTU4MjAyNDg3MywibmJmIjoxNTgyMDI0ODczLCJleHAiOjE1ODMzMjA4NzMsInN1YiI6IjI1NTU2MzkwLTBlZGMtMTFlYS04NGVlLWUxOTM1MzIwYTI1ZiIsInNjb3BlcyI6WyJ3bS1hcGk6YWRtaW4iLCJ3bS1hcGk6aW50ZWdyYXRpb24iXX0.lLcM2tGDkFPvT2_iE2mdmQqJB5M2kwlyVK_RmR-ikj-gt7TeSctCVD4d8Db1MB8UPl39Qz93BhxFR6jqSgC4rl92kHTNMMu0M0L3GY3um8EMdam5-P61RfypkDI--WXEXwN5cjtDnWtPaYFS1ukAkFkjDTc60q04GXda_lxWLRtKQfPWZTuV_4KyesJ0xcP-HatqUNgBCFrGDyZcr0LIdJ0Cpj-Exf0BzyjHSMsXs8uUMbZYgUKj3gENhAHJdctJHBVKS97RDR4wgPJbhO_BDlumXSHD9AxzNI9-396aLQ6TkP8UAosD0BJVzzOuFpjWGOjhGgl5lAkPcSrLxGVdhNPsRIlPo1Us3mqCfJpYpu0XJSYDZlZwtFtY67BnlyUuOGvAcmETtB0S2Abp4SHsl2_Vo1DHkezSnWZEFW0bsJ5EmwQkCBUT2tMc1MuiOfxPEOBv4gCrID4ea9pEPufxH8ZWMSWFizbgUwmmxUeB4_3uPpTzwuCDogB8ejr7cUxZMIh7iHIPG8uTkqFY9egQKxeoy_OiNWRvsT7ROF67BMsFUtKd_gjFGRC6G8kHZ3cFtZZTz1FHfJVZjb7VcaY3o3sep3wJ2yb8XMbyXvu3GiTVOG9DSYBMmwBLzyq-o66x6YoslDV018qNz7KGrPuO0WbPlSRew3lyAuW_pT-ZES8",
   "refresh_token":"def50200a2094ae53ad95b2876b79f7d802ae5078db1ca0555406406f9f3e111d374466fcd488ed373b8839bced16735dd7be2474d662793a21c577029191f8cb56714a8ba7cdf57d8bc1574e45e64fab60db908363d1d0c702feff7746a554df42772e61906fac7726b3c7b29b858aabb283fd2afb9e63590350a3b15ab5957c2267e75931b5a4fa9449e8671a847b793b7f7e4cdf811794f0bb156fe64a4d4c8a5773505d3f8801254dd03dd92fe24c91743fe5e844175a5a938d1b7b93f519e97604518017131958c15565c948db22bf3215262fc11360299df32a9f6215e037afa397f3ff928b3d560267b66ae63b2c1933ebb089b684038627b565b550bc8b2050a10e5a7ed933a2ba1bb98cf6b7259b93a818ec55d9936c89643f2efff494735a9c5cd91cd14b49501ffbddef750e8c01e9b32b83d60ee9758879b335767a6bbf9b3ba449f9ea0b8e6d5159f0c58554959798336ed646e4230aedc63bc0b4bf06c7549e0174fb56306248cc8fdbc19a686f76cd40f263c8cda6056dcd835dbdf79dbb38d9ed97c1d479aafbcb3de2d46d14ad1c8ec630720d43009ddc3a4d32293015a5b1345153fba689e3c037161bedea4b6b1d3edbf0fa1286b2a3f8918bad3ed96099887087e31"
}
```

#### grant_type `refresh_token`

>https://curity.io/resources/develop/oauth/oauth-refresh/

```sh
curl -XPOST 'https://api.aglaia.io/oauth/token' \
    -H 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjIzNzA1Yjg3YTA2YThiZDJiYTkwNjRlZjg5NzY0NzBjMTJhODQ4Mjg3MzNjOWVjNDA1NzMwNjYzYWFkOTdhMzE5NWVkOWRkMDkyNmNhOTJjIn0.eyJhdWQiOiIxM2RiMDY4MC1kNjZjLTExZTktOTgwMy0zMTRmNjA5MWZlZWUiLCJqdGkiOiIyMzcwNWI4N2EwNmE4YmQyYmE5MDY0ZWY4OTc2NDcwYzEyYTg0ODI4NzMzYzllYzQwNTczMDY2M2FhZDk3YTMxOTVlZDlkZDA5MjZjYTkyYyIsImlhdCI6MTU4MjAyNDg3MywibmJmIjoxNTgyMDI0ODczLCJleHAiOjE1ODMzMjA4NzMsInN1YiI6IjI1NTU2MzkwLTBlZGMtMTFlYS04NGVlLWUxOTM1MzIwYTI1ZiIsInNjb3BlcyI6WyJ3bS1hcGk6YWRtaW4iLCJ3bS1hcGk6aW50ZWdyYXRpb24iXX0.lLcM2tGDkFPvT2_iE2mdmQqJB5M2kwlyVK_RmR-ikj-gt7TeSctCVD4d8Db1MB8UPl39Qz93BhxFR6jqSgC4rl92kHTNMMu0M0L3GY3um8EMdam5-P61RfypkDI--WXEXwN5cjtDnWtPaYFS1ukAkFkjDTc60q04GXda_lxWLRtKQfPWZTuV_4KyesJ0xcP-HatqUNgBCFrGDyZcr0LIdJ0Cpj-Exf0BzyjHSMsXs8uUMbZYgUKj3gENhAHJdctJHBVKS97RDR4wgPJbhO_BDlumXSHD9AxzNI9-396aLQ6TkP8UAosD0BJVzzOuFpjWGOjhGgl5lAkPcSrLxGVdhNPsRIlPo1Us3mqCfJpYpu0XJSYDZlZwtFtY67BnlyUuOGvAcmETtB0S2Abp4SHsl2_Vo1DHkezSnWZEFW0bsJ5EmwQkCBUT2tMc1MuiOfxPEOBv4gCrID4ea9pEPufxH8ZWMSWFizbgUwmmxUeB4_3uPpTzwuCDogB8ejr7cUxZMIh7iHIPG8uTkqFY9egQKxeoy_OiNWRvsT7ROF67BMsFUtKd_gjFGRC6G8kHZ3cFtZZTz1FHfJVZjb7VcaY3o3sep3wJ2yb8XMbyXvu3GiTVOG9DSYBMmwBLzyq-o66x6YoslDV018qNz7KGrPuO0WbPlSRew3lyAuW_pT-ZES8' \
    -H 'content-type: application/x-www-form-urlencoded' \
    --data 'grant_type=refresh_token&refresh_token=def50200fae7423d7f6da8a50ab8ed3788c0736ae9872afcb58c4fb655acefd1ac7783c50e1809fcf5aa2d57f797e6ea4329a032cf8e527d2efba7ee27151dc9ea30d512ef05eeb27b9b7bd26f590f8f0a45891973d00dd1dbae1ab59fa170e2bad386be8324d34e93feed510c8e38c717d96ef75df77bc46b2639796652937a6a030232600c142ee87d4dbfaea1ea122c4f25c210312840ea9c0a2e683f65753369858a87530a9a137f4ac2d96a3a763c7089c1ee511e01d83d17f1baa959a475fa5c93be2170271965710decf9fb93a4b365557ebf4463010800f02476b37de4170765e8279e330ebd5143d49a338955214887eba4b2344bf8c36537ed2d533c8d8af3bd337fc7a2673933b6aae41c95dfa5f7440de6895a8be82260ec38ad185dfc2996de20706c69fa1d237f9bd0c3c4be0ab59c83b2481f1dc089a0e1ea0db054967443d3f4e2767b161ca9ea7eef96f2c695f4b08717c0a667e6a3b6f75446038a54281395cf0d74af66db3ad35a6b4106d9bea8865e80918040e1823088f1337ceb9880455872d8bb3c6fa9d295bf5d11d5573c89481ac95b41eef02afa1e11439480e97233ffc8ed44742751d2000a7fc340fb77a7bd1e4280c05468611dd0e35356810b32721069' \
    --compressed
```

```json
{
   "token_type":"Bearer",
   "expires_in":1296000,
   "access_token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjZmNzdlMzZhM2U5OWQxMjllODk3YTY3ZjI1MjQ0NTRjODA4MjQ3YzE3OGE5YmQwNGNkNDQzY2FjNDRkMjdkZTU1NDhlOGRkZDE5YThkZjA5In0.eyJhdWQiOiIxM2RiMDY4MC1kNjZjLTExZTktOTgwMy0zMTRmNjA5MWZlZWUiLCJqdGkiOiI2Zjc3ZTM2YTNlOTlkMTI5ZTg5N2E2N2YyNTI0NDU0YzgwODI0N2MxNzhhOWJkMDRjZDQ0M2NhYzQ0ZDI3ZGU1NTQ4ZThkZGQxOWE4ZGYwOSIsImlhdCI6MTU4MjAyOTE3MywibmJmIjoxNTgyMDI5MTczLCJleHAiOjE1ODMzMjUxNzMsInN1YiI6IjI1NTU2MzkwLTBlZGMtMTFlYS04NGVlLWUxOTM1MzIwYTI1ZiIsInNjb3BlcyI6WyJ3bS1hcGk6YWRtaW4iLCJ3bS1hcGk6aW50ZWdyYXRpb24iXX0.rfjkySqZ-41pKzbQ5L2FtQ5yXvJjqLgXKAYGUoJMi1lpLLIbebhbjdNLANqrfq2-dSEQBH4Om3-EKa004dhdsjezSBtOdw2cuOhksv327jQLakN6hQoL7HJE0ibFYFxojixjZN1Zczo-5c91QxeVysz2EI7UA2h-oaPgJeoNybIkiveiv8eAhRpForFXplB-Ps_RKGQ4WyuQihh08GnUeF4a8c07gnkHEOlhJy9XaIgC7s7Ifcz4njn5FiKcm74gOyA1Ue72Zs2JxLD3ANIHow9ETtTy8ptZWWacSr3wBSnd08auaa6y0MhUZD-bOww3lbvX4Wv1QplaBaDj7FVfWoqrXOy6T3gxP6NNgmboMGwSJdDa74booTgh1rXTrqmWU6t_4skH6892o39kGnFxK_n26h8n_9tSLlx3BRsT0UkWU-fBRAfg2NN_88UQm5HS1uKLRgozNKuPJMxL4rWk9HUunarBuS2eaESh8yzj5MNNV6CDjHS3FMSWK2UAvLMYvb0KYTeG4Xh6U6qFPKCa34fBtlrcmAHVozO45pSPVJqnfCL4L2tnAHS73W6aOLMzLtyejo__2cYhztTX-YsyDUTWkMNe0uWy5seZ1QziWQQSpFvjTuBvB8KJwpWe8sf39h9i1KuSfRIxlKcpxSfrNnR00UdE24BRqOWPtuEvYg0",
   "refresh_token":"def50200b410a2f87eefda9de2761664b52e0171b7a02442f4588d69f77b546662b9e19f0a5fa5afe0ffbada700ddebdab3090281b1b050e020480b0c15791ede9fe850b2e37e062775017f7f1e0ecaed4aaf2f9e9bcfd176b1cc5c4dc98fcbcd20c91fb01316069320c1b5ec579062ec2ad78a77270fa47ba4e3ced4e06861b7ab7b3416d384a93daac06869136a42c20623c75cee1007019763ed1f114cb2e6512738d0da11bfaf15137af45c3b95589a970572515ceeb6c3dd0974854a8bb5183a71bf0cf40c27c3147f24f3c2405109e115abc1fb176909dc67e9a6915a18ebcd3b82348b4d11b00fac70d4fe400dc4231fed49a69b619a99b113474c5a232154df597ee3caa98cc804268b8b6a723206c1e5fc34550cc5c7934cb5c123550c72a0cda02b69df4ad1596788825586a740dc845f600e8deb72a71e9f8c8c779be0defddc9a745d29c32fa43ced4f8fa8ba27c3148353ab1720731c02702b99c2a55b52ac0682e8093db97470c9d2bd0396b435a12f1f1bc185e6b3a8a208f2c2d937e2642ddef26e4a143cfcac5ce22a62bc01d9c75eff653dbcde749c7346e357f15337366ecf06ab977598b7c50ad4a12bab8475c7cdc008c50bcf950b63a5e42b6663aa62de38c1171"
}
```


#### grant_type `client_credentials`

>https://curity.io/resources/develop/oauth/oauth-client-credentials-flow/

```sh
curl -XPOST 'https://api.aglaia.io/oauth/token' \
    -H 'authorization: Basic MTNkYjA2ODAtZDY2Yy0xMWU5LTk4MDMtMzE0ZjYwOTFmZWVlOkltbGVVUUNCM0lORW9PNVVBU3AzY3FiM2ZBYkxVc2VJcm8wU0NxcGQ=' \
    -H 'content-type: application/x-www-form-urlencoded' \
    --data 'grant_type=client_credentials&scope=wm-api:admin+wm-api:integration' \
    --compressed
```

```json
{
   "token_type":"Bearer",
   "expires_in":1296000,
   "access_token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjdjMzY3NmZlMDRmZjRhNjBiMTIyZTJmODA3ZWMxNjEyMDhhN2YzNWI5MWY0MmFiYWZmYjMyODUwODhjZmM1ZGJjOGE0NjdjMTU0MWU0OTg4In0.eyJhdWQiOiIxM2RiMDY4MC1kNjZjLTExZTktOTgwMy0zMTRmNjA5MWZlZWUiLCJqdGkiOiI3YzM2NzZmZTA0ZmY0YTYwYjEyMmUyZjgwN2VjMTYxMjA4YTdmMzViOTFmNDJhYmFmZmIzMjg1MDg4Y2ZjNWRiYzhhNDY3YzE1NDFlNDk4OCIsImlhdCI6MTU4MjAyODgyMSwibmJmIjoxNTgyMDI4ODIxLCJleHAiOjE1ODMzMjQ4MjEsInN1YiI6IiIsInNjb3BlcyI6WyJ3bS1hcGk6YWRtaW4iLCJ3bS1hcGk6aW50ZWdyYXRpb24iXX0.umbHfxNq56Rz-DSoArH7oR8Gfu86C-CEzBPp_Yk9YlkWCuTx0uueMDhDJ-EXJG-3zeMW_zhgoCESrY9aCSBiG1HWCLstffznQ3k2IoKPwR2nqNX3B1E2JqNXED26DX-1Wi5CL60lPuCFROc6k0fn_0lWVXX6qGcjPqrWIKQxX4hJB6NyR9Qc9C9TQJQPotkNi8a7F56GLLTtgFw8XkxsAos09TvBNRB9evWPyi1rMfQ3EwERTw85MmhgUNEoQoqBsRRFsz1qAaGrMxpew9Q8PDpujAtfODoejF7aw7X8XeybuBtYPOUOIwSec5n0bgKTiON5t4yaIs-5zgYu9bVvm5Hd1dI0OkkSLyg3ol--cWSGEHySmKHYIBgFxI143_ealW6Xey_gSk8LbLu49Vtg9CJxHAAR3bpuxiVKwk7HHqfuQPNIN95u3brcWmzXPhTA2FYnLcHD54jcwlna8YaTQkD_KOrYvmPSWIzzRKM55HIBbYeHr_kSnNQFye-OcHnI-_K_Vd-xXyrDX91W4nLObfEkajYvCOzP7Jcwa5MyqTGY5tWK0FklwfEKVXTUixMA8exR9vknt8YwPl0SFoonDkZ4RzSbVtNXWPNeM3Ks6sSkozHPbmIj5XwQ0fICG4I4C3mbrvM6tx3adFIEKmauk9Vzr5xlA-uOQCr-FZLBG80"
}
```

#### grant_type `implicit`

>https://curity.io/resources/develop/oauth/oauth-implicit-flow/

redirect to uri for authentication

```sh
curl -XGET 'http://localhost:9099/oauth2/authorize?response_type=token&client_id=a9832dab-598c-11ea-a5a2-0242c0a8a000&redirect_uri=https:%2F%2Fwebhook.site%2F365c600d-ce97-471a-805e-6076eef7f9aa&scope=user:read%20user:write%20user:delete&state=teste_1'
```

callback with access_token

```json
{
   "token_type":"Bearer",
   "expires_in":1296000,
   "access_token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjZmNzdlMzZhM2U5OWQxMjllODk3YTY3ZjI1MjQ0NTRjODA4MjQ3YzE3OGE5YmQwNGNkNDQzY2FjNDRkMjdkZTU1NDhlOGRkZDE5YThkZjA5In0.eyJhdWQiOiIxM2RiMDY4MC1kNjZjLTExZTktOTgwMy0zMTRmNjA5MWZlZWUiLCJqdGkiOiI2Zjc3ZTM2YTNlOTlkMTI5ZTg5N2E2N2YyNTI0NDU0YzgwODI0N2MxNzhhOWJkMDRjZDQ0M2NhYzQ0ZDI3ZGU1NTQ4ZThkZGQxOWE4ZGYwOSIsImlhdCI6MTU4MjAyOTE3MywibmJmIjoxNTgyMDI5MTczLCJleHAiOjE1ODMzMjUxNzMsInN1YiI6IjI1NTU2MzkwLTBlZGMtMTFlYS04NGVlLWUxOTM1MzIwYTI1ZiIsInNjb3BlcyI6WyJ3bS1hcGk6YWRtaW4iLCJ3bS1hcGk6aW50ZWdyYXRpb24iXX0.rfjkySqZ-41pKzbQ5L2FtQ5yXvJjqLgXKAYGUoJMi1lpLLIbebhbjdNLANqrfq2-dSEQBH4Om3-EKa004dhdsjezSBtOdw2cuOhksv327jQLakN6hQoL7HJE0ibFYFxojixjZN1Zczo-5c91QxeVysz2EI7UA2h-oaPgJeoNybIkiveiv8eAhRpForFXplB-Ps_RKGQ4WyuQihh08GnUeF4a8c07gnkHEOlhJy9XaIgC7s7Ifcz4njn5FiKcm74gOyA1Ue72Zs2JxLD3ANIHow9ETtTy8ptZWWacSr3wBSnd08auaa6y0MhUZD-bOww3lbvX4Wv1QplaBaDj7FVfWoqrXOy6T3gxP6NNgmboMGwSJdDa74booTgh1rXTrqmWU6t_4skH6892o39kGnFxK_n26h8n_9tSLlx3BRsT0UkWU-fBRAfg2NN_88UQm5HS1uKLRgozNKuPJMxL4rWk9HUunarBuS2eaESh8yzj5MNNV6CDjHS3FMSWK2UAvLMYvb0KYTeG4Xh6U6qFPKCa34fBtlrcmAHVozO45pSPVJqnfCL4L2tnAHS73W6aOLMzLtyejo__2cYhztTX-YsyDUTWkMNe0uWy5seZ1QziWQQSpFvjTuBvB8KJwpWe8sf39h9i1KuSfRIxlKcpxSfrNnR00UdE24BRqOWPtuEvYg0",
   "refresh_token":"def50200b410a2f87eefda9de2761664b52e0171b7a02442f4588d69f77b546662b9e19f0a5fa5afe0ffbada700ddebdab3090281b1b050e020480b0c15791ede9fe850b2e37e062775017f7f1e0ecaed4aaf2f9e9bcfd176b1cc5c4dc98fcbcd20c91fb01316069320c1b5ec579062ec2ad78a77270fa47ba4e3ced4e06861b7ab7b3416d384a93daac06869136a42c20623c75cee1007019763ed1f114cb2e6512738d0da11bfaf15137af45c3b95589a970572515ceeb6c3dd0974854a8bb5183a71bf0cf40c27c3147f24f3c2405109e115abc1fb176909dc67e9a6915a18ebcd3b82348b4d11b00fac70d4fe400dc4231fed49a69b619a99b113474c5a232154df597ee3caa98cc804268b8b6a723206c1e5fc34550cc5c7934cb5c123550c72a0cda02b69df4ad1596788825586a740dc845f600e8deb72a71e9f8c8c779be0defddc9a745d29c32fa43ced4f8fa8ba27c3148353ab1720731c02702b99c2a55b52ac0682e8093db97470c9d2bd0396b435a12f1f1bc185e6b3a8a208f2c2d937e2642ddef26e4a143cfcac5ce22a62bc01d9c75eff653dbcde749c7346e357f15337366ecf06ab977598b7c50ad4a12bab8475c7cdc008c50bcf950b63a5e42b6663aa62de38c1171"
}
```

#### grant_type `authorization_code`

>https://curity.io/resources/develop/oauth/oauth-code-flow/

redirect to uri for authentication

```sh
curl -XGET 'http://localhost:9091/oauth2/authorize?response_type=code&client_id=a9832dab-598c-11ea-a5a2-0242c0a8a000&scope=user:read%20user:write%20user:delete&state=teste_123&redirect_uri=https://webhook.site/365c600d-ce97-471a-805e-6076eef7f9aa'
```

http status 302 redirect with `code`

```http
https://webhook.site/5bb5b879-8e43-4322-9fc0-50c30325516b?code=def502007131aedd71cdc08dab9dfccb271600e94a8297f4bcdd056907d91a52f4c48bf66ad732ef69dc93362c306d95abc928b73c14f57a084e467b10fd04c76fd3f6e50cd32d834f4d051fba26d4e388ad09ab125a731774b1c52e2020a0107b127b8d5e86bbef50396e5709d64c4366076093613ab4c41b5e1a5370d60eb7e1890d48cbaef52b95c23f8b6e790c8019a27155620cc536348c7e4d8fe943362340754c891c6ce7197735df48204ff165190a117e8a3b37504d3389d7ae57b4e612f05b2fd491c90fbbe5b114ff25328bd57abd64c54107af49933ed3370487dae9ce6238b124ff768d9d08a3408f981f975c8339e5246d1d930800f124228b64f64f9cfaf3053fc367aa5ee0f44495ecee0d3b9fea6d646e9c69cdf85efaf01581c0c18d111f07901788bbc9c180de30913cd0b10d09333f5da5cf83656449702ba038c302246fc7020250d035f464d21d5342462b68099bcaab7e8f687aa2c62a32d4dc1a3e476c594b8e0676d2af0573842a8ff7469f74ea0e5a1e16fdc03ed91da067b4fe4737a884a253248928b500ccb2f6a1d1d0a05e56062b5da70ffb57b7a99effb3733e6c7a1ae1f8142108576af86033955cd76df9dfd53c9b7cfa20648fc1d9ef626bef0e3f25c0a4f2dc&state=xcoiv98y2kd22vusuye3kch
```

request for obtain access_token

>Observation, send client_secred if authentication is not use basic_authentication
```sh
curl -XPOST 'http://localhost:9091/oauth2/token' \
    -H 'authorization: Basic MTNkYjA2ODAtZDY2Yy0xMWU5LTk4MDMtMzE0ZjYwOTFmZWVlOkltbGVVUUNCM0lORW9PNVVBU3AzY3FiM2ZBYkxVc2VJcm8wU0NxcGQ=' \
    -H 'content-type: application/x-www-form-urlencoded' \
    --data 'grant_type=authorization_code&client_id=a9832dab-598c-11ea-a5a2-0242c0a8a000&client_secret=8zd9ULma6xNN1wbR7h8er7z7qbERULsjCqD2pzT5&redirect_uri=https://webhook.site/5bb5b879-8e43-4322-9fc0-50c30325516b&code=def502007131aedd71cdc08dab9dfccb271600e94a8297f4bcdd056907d91a52f4c48bf66ad732ef69dc93362c306d95abc928b73c14f57a084e467b10fd04c76fd3f6e50cd32d834f4d051fba26d4e388ad09ab125a731774b1c52e2020a0107b127b8d5e86bbef50396e5709d64c4366076093613ab4c41b5e1a5370d60eb7e1890d48cbaef52b95c23f8b6e790c8019a27155620cc536348c7e4d8fe943362340754c891c6ce7197735df48204ff165190a117e8a3b37504d3389d7ae57b4e612f05b2fd491c90fbbe5b114ff25328bd57abd64c54107af49933ed3370487dae9ce6238b124ff768d9d08a3408f981f975c8339e5246d1d930800f124228b64f64f9cfaf3053fc367aa5ee0f44495ecee0d3b9fea6d646e9c69cdf85efaf01581c0c18d111f07901788bbc9c180de30913cd0b10d09333f5da5cf83656449702ba038c302246fc7020250d035f464d21d5342462b68099bcaab7e8f687aa2c62a32d4dc1a3e476c594b8e0676d2af0573842a8ff7469f74ea0e5a1e16fdc03ed91da067b4fe4737a884a253248928b500ccb2f6a1d1d0a05e56062b5da70ffb57b7a99effb3733e6c7a1ae1f8142108576af86033955cd76df9dfd53c9b7cfa20648fc1d9ef626bef0e3f25c0a4f2dc' \
    --compressed
```

```json
{
   "token_type":"Bearer",
   "expires_in":1296000,
   "access_token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjZmNzdlMzZhM2U5OWQxMjllODk3YTY3ZjI1MjQ0NTRjODA4MjQ3YzE3OGE5YmQwNGNkNDQzY2FjNDRkMjdkZTU1NDhlOGRkZDE5YThkZjA5In0.eyJhdWQiOiIxM2RiMDY4MC1kNjZjLTExZTktOTgwMy0zMTRmNjA5MWZlZWUiLCJqdGkiOiI2Zjc3ZTM2YTNlOTlkMTI5ZTg5N2E2N2YyNTI0NDU0YzgwODI0N2MxNzhhOWJkMDRjZDQ0M2NhYzQ0ZDI3ZGU1NTQ4ZThkZGQxOWE4ZGYwOSIsImlhdCI6MTU4MjAyOTE3MywibmJmIjoxNTgyMDI5MTczLCJleHAiOjE1ODMzMjUxNzMsInN1YiI6IjI1NTU2MzkwLTBlZGMtMTFlYS04NGVlLWUxOTM1MzIwYTI1ZiIsInNjb3BlcyI6WyJ3bS1hcGk6YWRtaW4iLCJ3bS1hcGk6aW50ZWdyYXRpb24iXX0.rfjkySqZ-41pKzbQ5L2FtQ5yXvJjqLgXKAYGUoJMi1lpLLIbebhbjdNLANqrfq2-dSEQBH4Om3-EKa004dhdsjezSBtOdw2cuOhksv327jQLakN6hQoL7HJE0ibFYFxojixjZN1Zczo-5c91QxeVysz2EI7UA2h-oaPgJeoNybIkiveiv8eAhRpForFXplB-Ps_RKGQ4WyuQihh08GnUeF4a8c07gnkHEOlhJy9XaIgC7s7Ifcz4njn5FiKcm74gOyA1Ue72Zs2JxLD3ANIHow9ETtTy8ptZWWacSr3wBSnd08auaa6y0MhUZD-bOww3lbvX4Wv1QplaBaDj7FVfWoqrXOy6T3gxP6NNgmboMGwSJdDa74booTgh1rXTrqmWU6t_4skH6892o39kGnFxK_n26h8n_9tSLlx3BRsT0UkWU-fBRAfg2NN_88UQm5HS1uKLRgozNKuPJMxL4rWk9HUunarBuS2eaESh8yzj5MNNV6CDjHS3FMSWK2UAvLMYvb0KYTeG4Xh6U6qFPKCa34fBtlrcmAHVozO45pSPVJqnfCL4L2tnAHS73W6aOLMzLtyejo__2cYhztTX-YsyDUTWkMNe0uWy5seZ1QziWQQSpFvjTuBvB8KJwpWe8sf39h9i1KuSfRIxlKcpxSfrNnR00UdE24BRqOWPtuEvYg0",
   "refresh_token":"def50200b410a2f87eefda9de2761664b52e0171b7a02442f4588d69f77b546662b9e19f0a5fa5afe0ffbada700ddebdab3090281b1b050e020480b0c15791ede9fe850b2e37e062775017f7f1e0ecaed4aaf2f9e9bcfd176b1cc5c4dc98fcbcd20c91fb01316069320c1b5ec579062ec2ad78a77270fa47ba4e3ced4e06861b7ab7b3416d384a93daac06869136a42c20623c75cee1007019763ed1f114cb2e6512738d0da11bfaf15137af45c3b95589a970572515ceeb6c3dd0974854a8bb5183a71bf0cf40c27c3147f24f3c2405109e115abc1fb176909dc67e9a6915a18ebcd3b82348b4d11b00fac70d4fe400dc4231fed49a69b619a99b113474c5a232154df597ee3caa98cc804268b8b6a723206c1e5fc34550cc5c7934cb5c123550c72a0cda02b69df4ad1596788825586a740dc845f600e8deb72a71e9f8c8c779be0defddc9a745d29c32fa43ced4f8fa8ba27c3148353ab1720731c02702b99c2a55b52ac0682e8093db97470c9d2bd0396b435a12f1f1bc185e6b3a8a208f2c2d937e2642ddef26e4a143cfcac5ce22a62bc01d9c75eff653dbcde749c7346e357f15337366ecf06ab977598b7c50ad4a12bab8475c7cdc008c50bcf950b63a5e42b6663aa62de38c1171"
}
```


### Database structure

- oauth_users
	- id
    - name
    - last_name
    - email
    - username
    - email
    - password
    - birthday
    - activated
    - created_at
    - updated_at
    - expires_at
    - deleted_at
- oauth_roles
	- id
    - name
- oauth_user_roles
	- user_id
    - role_id
- oauth_clients
	- id
    - user_id
    - name
    - scopes
    - redirect
    - revoked
    - created_at
    - updated_at
- oauth_access_tokens
	- id
    - user_id
    - client_id
    - access_token
    - scopes
    - revoked
    - expires_at
- oauth_refresh_tokens
	- id
    - access_token_id
    - refresh_token
    - revoked
    - created_at
    - expires_at
- oauth_auth_codes
	- id
    - user_id
    - client_id
    - code
    - scopes
    - revoked
    - created_at
    - expires_at