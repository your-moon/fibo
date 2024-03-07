import { Card, CardBody, CardHeader } from "@nextui-org/react";

interface BigCardProps {
  title: string;
  value: number | string | undefined;
  valueColor?: string;
}

export const BigCard = (props: BigCardProps) => {
  return (
    <Card
      radius="lg"
      className={`mr-6  flex-auto min-w-[320px] max-w-[460px] min-h-[250px]`}
    >
      <CardHeader className=" ml-10 mt-6 justify-between">
        <h3 className="font-bold">{props.title}</h3>
      </CardHeader>
      <CardBody className="flex justify-center items-center">
        <p className={`text-6xl font-bold py-5 ${props.valueColor}`}>
          {props.value}
        </p>
      </CardBody>
    </Card>
  );
};
