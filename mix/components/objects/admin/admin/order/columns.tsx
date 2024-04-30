"use client"
 
import { Button } from "@/components/ui/button"
import { ColumnDef } from "@tanstack/react-table"
 
// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.
export type Order = {
  id: string
  user_id: string
  total_price: number
  status: string
  created_at: string
  view?: () => void
  
}

export const columns: ColumnDef<Order>[] = [
  {
    accessorKey: "id",
    header: "ID",
  },
  {
    accessorKey: "user_id",
    header: "UserId",
  },
  {
    accessorKey: "total_price",
    header: "TotalPrice",
  },
  {
    accessorKey: "created_at",
    header: "CreatedAt",
  },
  {
    accessorKey: "Actions",
    header: "Actions",
    cell: ({ row }) => (
      <div>
        <Button  onClick={row.original.view}  className="text-sm px-2 py-1 mr-2 bg-blue-500 hover:bg-blue-700 text-white rounded" >View</Button>
      </div>
    ),
  },

  
]