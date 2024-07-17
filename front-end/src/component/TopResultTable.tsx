export const TopResultTable = ({
  result,
  index,
  handleNavigate,
}: {
  result: SearchResponse;
  index: number;
  handleNavigate: (type: string, result: SearchResponse) => void;
}) => {
  return (
    <div
      className="topTable"
      key={result.song.songId}
      onClick={() => {
        handleNavigate("song", result);
      }}
    >
      <div className="title">
        <p>{index + 1}. </p>
        <img src={result.song.album.banner} alt="Song Cover" />
        <div>
          <h3>{result.song.title}</h3>
          <p>{result.song.artist.user.username}</p>
        </div>
      </div>
      <p>
        {Math.floor(result.song.duration / 60)}:
        {Math.floor(result.song.duration % 60)
          .toString()
          .padStart(2, "0")}
      </p>
    </div>
  );
};
