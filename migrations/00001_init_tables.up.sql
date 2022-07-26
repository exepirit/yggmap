-- nodes definition

CREATE TABLE nodes (
	public_key BLOB NOT NULL,
	coordinates TEXT NOT NULL,
	additional_info BLOB,
	CONSTRAINT nodes_pk PRIMARY KEY (public_key)
);

-- peer_links definition

CREATE TABLE peer_links (
	key1 BLOB NOT NULL,
	key2 BLOB NOT NULL,
	CONSTRAINT peer_links_FK_1 FOREIGN KEY (key1) REFERENCES nodes(public_key),
	CONSTRAINT peer_links_FK_2 FOREIGN KEY (key2) REFERENCES nodes(public_key)
);