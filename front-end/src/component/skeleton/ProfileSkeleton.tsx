import { Shimmer } from "./Shimmer.tsx";
import { Skeleton } from "./Skeleton.tsx";

export const ProfileSkeleton = () => {
  return (
    <div className={"skeletonWrapper"}>
      <div className={"skeleton"}>
        <Skeleton type={"profile"} />
        <Skeleton type={"title"} />
        <Skeleton type={"text"} />
      </div>
      <Shimmer />
    </div>
  );
};
