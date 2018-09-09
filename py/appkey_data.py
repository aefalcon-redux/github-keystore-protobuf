import appkey_pb2
default_links_pb = b'\n\tindex.pbf\x12\x0fapp/{AppId}.pbf\x1a\x1dkey/{AppId}/{Fingerprint}.bin""key/{AppId}/{Fingerprint}-meta.pbf'
default_links = appkey_pb2.Links()
default_links.ParseFromString(default_links_pb)
