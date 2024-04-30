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
  
]