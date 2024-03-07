import { Spinner } from "@nextui-org/react";

const Loading = () => {
  return (
    <div className=" my-20 flex flex-col items-center justify-center">
      <Spinner size="lg" color="success" />
    </div>
  );
};

export default Loading;
