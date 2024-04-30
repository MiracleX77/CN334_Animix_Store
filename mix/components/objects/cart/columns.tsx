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
  count: number
  onRemove: () => void
}
 
export const columns: ColumnDef<Product>[] = [
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
    accessorKey: "count",
    header: "Count",
  },
  {
    accessorKey: "price",
    header: "Price",
  },
  {
    accessorKey: "Actions",
    header: "Actions",
    cell: ({ row }) => (
      <div>
        <Button onClick={row.original.onRemove} className="text-sm px-2 py-1 mr-2 bg-blue-500 hover:bg-blue-700 text-white rounded" >Delete</Button>
      </div>
    ),
  },

  
]