import appkey_pb2
default_links_pb = b'\n\tindex.pbf\x12\x0bapp/{AppId}\x1a\x1dapp/{AppId}/{Fingerprint}.bin""app/{AppId}/{Fingerprint}-meta.pbf'
default_links = appkey_pb2.Links()
default_links.ParseFromString(default_links_pb)
