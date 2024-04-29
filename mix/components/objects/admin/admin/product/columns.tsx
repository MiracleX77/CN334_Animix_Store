"use client"
 
import { Button } from "@/components/ui/button"
import { ColumnDef } from "@tanstack/react-table"
 
// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.
export type Product = {
  id: string
  img_url: string
  name: string
  price: number
  stock: number
  status: "Active" | "Inactive"
  edit: () => void
  delete: () => void
}
 
export const columns: ColumnDef<Product>[] = [
  {
    accessorKey: "id",
    header: "ID",
  },
  {
    accessorKey: "img",
    header: "Image",
    cell: ({ row }) => (
      <img src={row.original.img_url} alt="Product" width={100} height={100} />
    ),
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "price",
    header: "Price",
  },
  {
    accessorKey: "stock",
    header: "Stock",
  },
  {
    accessorKey: "status",
    header: "Status",
  },
  {
    accessorKey: "Actions",
    header: "Actions",
    cell: ({ row }) => (
      <div>
        <Button  onClick={row.original.edit} className="bg-sec ring-2 px-8 rounded-xl text-white text-sm px-2 py-1 mr-2 bg-blue-500 hover:bg-blue-700 text-white rounded">Edit</Button>
        <Button  onClick={row.original.delete}  className="text-sm px-2 py-1 mr-2 bg-blue-500 hover:bg-blue-700 text-white rounded" >Delete</Button>
      </div>
    ),
  },

  
]