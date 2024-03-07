import { BACKEND_URL } from "@/app/provider";
import { Select, SelectItem, Selection, Spinner } from "@nextui-org/react";
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";
import { Dispatch, SetStateAction, useEffect, useState } from "react";

export interface CatResponse {
  status: string;
  message: string;
  data: SingleCat[];
}

export interface SingleCat {
  Id: number;
  Name: string;
  CreatedAt: string;
  UpdatedAt: string;
}

interface CategoryCompProps {
  categoryId: string;
  setCategoryId: Dispatch<SetStateAction<string>>;
}

export default function CategoryComp({
  categoryId,
  setCategoryId,
}: CategoryCompProps) {
  return (
    <QueryClientProvider client={new QueryClient()}>
      <Categories categoryId={categoryId} setCategoryId={setCategoryId} />
    </QueryClientProvider>
  );
}

function Categories({ categoryId, setCategoryId }: CategoryCompProps) {
  const [categories, setCategories] = useState<SingleCat[]>([]);
  const { isPending, error, data } = useQuery({
    queryKey: ["catsX"],
    queryFn: () =>
      fetch(`${BACKEND_URL}/categories`, {}).then(
        (res) => res.json() as Promise<CatResponse>,
      ),
  });

  useEffect(() => {
    if (data) {
      setCategories(data.data);
    }
  }, [data]);

  if (isPending) return <Spinner color="success" className="mx-5" />;

  if (error)
    return (
      <Select className="min-w-[200px] max-w-250px mx-4" variant="bordered">
        <SelectItem key="0" value="0">
          Error
        </SelectItem>
      </Select>
    );

  console.log(categories);
  return (
    <div className="min-w-[200px] max-w-250px mx-4">
      <Select
        variant="bordered"
        label="Select a category"
        selectedKeys={[categoryId]}
        onChange={(e) => setCategoryId(e.target.value)}
      >
        {categories.map((cat) => (
          <SelectItem key={cat.Id} value={cat.Name}>
            {cat.Name}
          </SelectItem>
        ))}
      </Select>
    </div>
  );
}
