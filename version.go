/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

const (
    VERSION_COPYRIGHT = "Copyright David Thorpe 2016 All Rights Reserved"
    VERSION_AUTHOR    = "David Thorpe <djt@mutablelogic.com>"
    VERSION_URL       = "https://github.com/djthorpe/ytapi"
    VERSION_TAG       = "alpha-002-20-g0f3cf81"
    VERSION_BRANCH    = "master"
    VERSION_HASH      = "0f3cf817919310a3a5730869850a1bd0b91502d3"
    VERSION_DATE      = "Mon Jun 20 17:09:47 UTC 2016"
    VERSION_GOVERSION = "go version go1.6.2 darwin/amd64"
    VERSION_CLIENT_SECRET = "eyJpbnN0YWxsZWQiOnsiY2xpZW50X2lkIjoiMTA0NTg2NDA0NzYwLW5kM3JucTcxdGV1MzFlcHZ0NmkzbG9oaTI2MGo3YWJzLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwicHJvamVjdF9pZCI6Inl0YXBpLTEyODMiLCJhdXRoX3VyaSI6Imh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbS9vL29hdXRoMi9hdXRoIiwidG9rZW5fdXJpIjoiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tL28vb2F1dGgyL3Rva2VuIiwiYXV0aF9wcm92aWRlcl94NTA5X2NlcnRfdXJsIjoiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vb2F1dGgyL3YxL2NlcnRzIiwiY2xpZW50X3NlY3JldCI6IlNrVVJoclVOd3BtRGN4dmhSRlZoS2Z2cyIsInJlZGlyZWN0X3VyaXMiOlsidXJuOmlldGY6d2c6b2F1dGg6Mi4wOm9vYiIsImh0dHA6Ly9sb2NhbGhvc3QiXX19"
    VERSION_SERVICE_ACCOUNT = "ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAieXRhcGktMTI4MyIsCiAgInByaXZhdGVfa2V5X2lkIjogImVlNGQ1ZTkxYjM0MDc4YTA5ZGMyYWMzMzdjZDg5MTcyZmNlMjk5YmMiLAogICJwcml2YXRlX2tleSI6ICItLS0tLUJFR0lOIFBSSVZBVEUgS0VZLS0tLS1cbk1JSUV2d0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktrd2dnU2xBZ0VBQW9JQkFRQ1ZhMUp1RXNWZ3l2cjRcbkJjczZwNnBCcEh6bGVXa3Q3bXB2eExVeDJ0VkdJYml4K1NTV2RBZ1pKZW0rc1d6bmprTWdzUDVQdWVFTHB4R3Zcbng2UmtaUnZoa2ZOR1NSSHpRdFVWZ21MeTg3SlNFa2V1QVhQbW9rRGErNXQzOGtMQm1NSzFqaXRFMko2YlpOTWlcbmRCR3pKSE4yL3F1Qkd3RTBjb2syTElOd3p1ZEIrVWRvUjNCc3V6ZW91cTdJdmNoSGl5Q2VCMXpsamx6c0NhbVVcblpKVzhDMXM1emYzR1g1RE5sT2xGc2FEQTA3S3BFRWdBYjFXR2tFbmZjNlU1dGNTM2kwRmV1TUx1Z3c1eThiN1JcbkdEdjRndWo1UElCVEVIUWlMZy82TUFRV09HUERNYmZKbkNPZjJQWkVZdmZHcXJsMTByc0hYcTdKd24vY2NCWGZcbkVHclk3ejd2QWdNQkFBRUNnZ0VBU0dtZnBDckNyenNQejcxM0ZYRUdLUlBmVjVLMGI0Y1kzQnZYdUZhV2Y0dCtcbkovcFhTZmJucWRCZTdJU3VBN1pGUUFXbTZkcHFmeXd6ZlIzaXhLamZqejBkU2hYRmRrQVFab0V2RlJFWGZZRFFcbmJqclVBWlRtbTNuT2laM0Z6bThMQmRQbGtsTHM0dUJGRm5CcjE0YTEwcGh2WTMxRDR6Z1MyVWRQV0E5ZHNmR3VcbjlHblp0TFNTbTk4am9uT0MyNFJoYXQxeTErOThNNi8ycGZHWU44eFdNRnk2YVYyVVhNR3V1RW9TZW9SRkdXT0RcbkMybFUzYTJEUUVKaGZVTk9rVWd3TENuYU9TTEpBK3RtQnhDaFBDT0V5SEdnZnhlNnJOM2VGbVVxbE9LOFlmZmNcbmtHSDFnM3krQnVIcVRNNmh4ZTFRTmpoMGw1aUlLNGI2b0szL0lQOTNnUUtCZ1FEZldtQ1ZZbnlwZWpwblQwL21cbndsYzFSKzBYNXQwNjBXNkZqK1VHNGNnMGNsRFliSDk1ZnZ5MlZNMVY0R0kxMEFnSHkvVnpRd3prOUZmb2pkejVcbnVPVnh1QndFOTEzMFNkcTFwTDgzakxrVDFCcVIydWZxSC81OHM0SXh6QzUrL0FueG1STC93NGhzcnlkTGFsSnNcblhnY1FKM3dkdHJzQ2VMdEw1dkVxVnhqSkx3S0JnUUNyUW14eHhQTFYyVytPNmtsMlpGajI2S2Q4anc0MkY0SnpcbkdORHFsekp0T2xLOEpncFFCQSsrZUtqUmNmQ2o4MWdHOURtb2JvVDVmbStHV1REdXZybG0zWXJXVk54ZG9ITThcbnBGV3BiRnJIb042U3Fud0RMRGNrNmVsLzFUeFpMWHhXWnBWTHN0V3lJd2RTSVFmTW05NkxHTkVnUFR0QkdVVGhcblAvVWxwSC8yUVFLQmdRQ3M4NUx1bWtEbUMxdkdPOGduYkQyVGFJclRnMVcwQWhDUEpBbTQ5b1RNOStoYWJrbVhcbmp4bXBKcW1KNGlLcWdOaDNFUnlUa0V3ZnpvVno4c1kyS1ZLaWFHcHAvaEttNDdtVUJtcDdYV0REVmU3dEdwR1NcbjNndDJESnE0Yk45a1BVNkY0bjZ2eG4yUmkxL21LZ0tCSDU1Z2gvSzRINU5ZS3B6OS9XcXAycCs2eVFLQmdRQ0NcbkpORWFSK3V0L1RXY3FvZzNFeHpVVFlyeHloaW5uSWdDaFVwRjY2NWplUjNrWXhnVkpIRCtjdHovaFJHMDFia0FcblY5bjFHRCt2Qlpra3hHUytHV3YxTXNoT0JtRU1PV0U1SlR6L3FocjJkMWFNcUVmamRWZFdrdmRoc0o3U3VoRXhcbjYzVHErTFVTVVRKQUJGL1Q2bU5jMnZnWkFzOGpzbHM0dll6REZWWTB3UUtCZ1FERFFtTllNeUtzcnh6Y2NuMHRcbms0UVgyV2dGZlZMWURnVXRmV25QSHl6NDE2Ti9VSUFjVUVWcEVWNTY5eGJwT2tzTjFmblg5UExRRW11S1IwTmFcbmlDWk9KZGNtMHdGZ25JTkYwWjNWeW1SdHRxVjlZWHMrSE9rSDlnbW1rSE1sZ0NzblNGYjAzQkQrc1ZKQ2JHeDJcbkVldFRBY25qVG50L2V2MldEM2czNTZGYW9RPT1cbi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS1cbiIsCiAgImNsaWVudF9lbWFpbCI6ICJ5dGFwaS1zZXJ2aWNlLWFjY291bnRAeXRhcGktMTI4My5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsCiAgImNsaWVudF9pZCI6ICIxMTMzMTE4ODEwNzM0NDAzMTI3NTgiLAogICJhdXRoX3VyaSI6ICJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20vby9vYXV0aDIvYXV0aCIsCiAgInRva2VuX3VyaSI6ICJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20vby9vYXV0aDIvdG9rZW4iLAogICJhdXRoX3Byb3ZpZGVyX3g1MDlfY2VydF91cmwiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vb2F1dGgyL3YxL2NlcnRzIiwKICAiY2xpZW50X3g1MDlfY2VydF91cmwiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vcm9ib3QvdjEvbWV0YWRhdGEveDUwOS95dGFwaS1zZXJ2aWNlLWFjY291bnQlNDB5dGFwaS0xMjgzLmlhbS5nc2VydmljZWFjY291bnQuY29tIgp9Cg=="
)

