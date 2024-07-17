interface Song {
  songId: string;
  title: string;
  artistId: string;
  albumId: string;
  releaseDate: string;
  duration: number;
  file: string;
  play: Play[];
  artist: Artist;
  album: Album;
  songAudio: HTMLAudioElement | null;
}
