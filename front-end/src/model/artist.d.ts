interface Artist {
  artistId: string;
  userId: string;
  description: string;
  banner: string;
  verifiedAt: string;
  user: User;
}

interface VerifyArtist {
  artist: Artist;
  follower: number;
  following: number;
}
