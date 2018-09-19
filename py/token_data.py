import appkey_pb2
default_links_pb = b'\n\x15app/{AppId}/token.pbf\x12%install/{AppId}/{InstallId}/token.pbf'
default_links = appkey_pb2.Links()
default_links.ParseFromString(default_links_pb)