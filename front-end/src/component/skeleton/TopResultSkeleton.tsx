import { Shimmer } from "./Shimmer.tsx";
import { Skeleton } from "./Skeleton.tsx";

export const TopResultSkeleton = () => {
  return (
    <div className={"skeletonTopWrapper"}>
      <div className={"skeleton"}>
        <Skeleton type={"title"} />
        <Skeleton type={"image"} />
        <Skeleton type={"title"} />
        <Skeleton type={"text"} />
      </div>
      <Shimmer />
    </div>
  );
};
