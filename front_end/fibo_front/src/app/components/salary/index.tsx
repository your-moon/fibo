import { Card, CardBody, CardHeader } from "@nextui-org/react";

export const Salary = () => {
  return (
    <Card
      radius="lg"
      className="mx-6 min-w-[380px] max-w-[460px] min-h-[250px]"
    >
      <CardHeader className=" ml-10 mt-6 justify-between">
        <h3 className="font-bold">Ongoing Salary</h3>
      </CardHeader>
      <CardBody className="flex justify-center items-center">
        <p className="text-6xl font-bold text-emerald-400 py-5">1223$</p>
      </CardBody>
    </Card>
  );
};
