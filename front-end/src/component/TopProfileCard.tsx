export const TopProfileCard = ({
  result,
  handleNavigate,
}: {
  result: SearchResponse;
  handleNavigate: (type: string, result: SearchResponse) => void;
}) => {
  return (
    <div
      className={"card"}
      onClick={() => {
        handleNavigate("artist", result);
      }}
    >
      <div className={"cardImage"}>
        <img
          src={
            result.song.artist.user.avatar
              ? result.song.artist.user.avatar
              : "/assets/download (6).png"
          }
          alt={"placeholder"}
          className={"profilePic"}
        />
      </div>
      <div className={"cardContent"}>
        <h3>{result.song.artist.user.username}</h3>
        <p>{result.song.artist.user.role}</p>
      </div>
    </div>
  );
};
