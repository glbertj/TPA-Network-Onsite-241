import { Shimmer } from "./Shimmer.tsx";
import { Skeleton } from "./Skeleton.tsx";

export const TableSkeleton = () => {
  return (
    <div className={"skeletonTableWrapper"}>
      <div className={"skeletonTable"}>
        <Skeleton type={"tableImage"} />
        <Skeleton type={"title"} />
        <Skeleton type={"text"} />
      </div>
      <Shimmer />
    </div>
  );
};
