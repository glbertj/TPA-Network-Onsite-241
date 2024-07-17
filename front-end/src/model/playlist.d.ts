interface Playlist {
    playlistId: string;
    userId : string;
    title: string;
    description: string;
    image: string;
    user : User;
    playlistDetails : PlaylistDetail[];
}

interface PlaylistDetail{
    playlistDetailId: string;
    playlistId: string;
    songId: string;
    song : Song;
    dateAdded : string;
}