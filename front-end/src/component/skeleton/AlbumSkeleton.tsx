import { Shimmer } from "./Shimmer.tsx";
import { Skeleton } from "./Skeleton.tsx";

export const AlbumSkeleton = () => {
  return (
    <div className={"skeletonWrapper"}>
      <div className={"skeleton"}>
        <Skeleton type={"image"} />
        <Skeleton type={"title"} />
        <Skeleton type={"text"} />
      </div>
      <Shimmer />
    </div>
  );
};
