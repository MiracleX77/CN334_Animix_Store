import ProductTable from "@/components/objects/admin/admin/product/ProductTable";
import ButtonAdd from "@/components/objects/admin/admin/product/ButtonAdd";




export default async function DashboardProductPage() {
  

  return (
    <div className="container mx-auto py-10">
      <div className="flex justify-between items-center mb-4">
        <div>
          <h1 className="text-2xl font-bold">Products</h1>
          <h2 className="text-lg">Manage your products</h2>
        </div>
        <ButtonAdd />
      </div>
      <ProductTable />
    </div>
  )
}