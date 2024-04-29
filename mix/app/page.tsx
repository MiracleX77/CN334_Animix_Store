import ListProductCard from "@/components/objects/ListProductCard";
import { ProductCard } from "@/components/objects/ProductCard";
import HomeCover from "@/components/objects/home/HomeCover";
import { List } from "lucide-react";

export default function Home() {

  return (
    <>
      <div className="mx-auto w-[80%] px-2 md:px-4 lg:px-6">
        <HomeCover />
        <h1 className="text-2xl font-bold">Best Seller</h1>
        <div className="mt-4 p-2 md:p-4 lg:p-6 ring-4 rounded-xl">

          <ListProductCard type='all' />
        </div>
      </div>
    </>
  );
}
