'use client'
import React, { useState, useEffect } from 'react';
import { DataTable } from './data-table';
import { getProducts,deleteProductById } from "@/apis/services/productService"
import { AuthProvider } from "@/utils/clientAuthProvider"
import { Product ,columns} from './columns';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { useRouter } from "next/navigation"
import CardEdit from './edit/CardEdit';

import { set } from 'date-fns';


function ProductTable() {

    const router = useRouter();
    const token = AuthProvider.getAccessToken()

    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertOnCancel, setAlertOnCancel] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');
    const [alertCancelBottom, setAlertCancelBottom] = useState(false);

    const [products, setProducts] = useState([]);

    useEffect(() => {
        loadProducts();
    }, []);

    const loadProducts = async () => {

        const fetchedProducts = await getProducts();
        // Attach edit and delete methods to each product
        const product = fetchedProducts.data.data
        product.map((product: Product) => {
            product.edit = () => router.push(`/dashboard/product/edit/${product.id}`);
            product.delete = () => handleDelete(product.id);
        });
        setProducts(product);
    };
    const deleteProduct = async (id:string) => {
        console.log("Delete product with id:", id);
        const res = await deleteProductById(id,token||'');
        if(res.status === 200){
            setAlertTitle("Delete Product");
            setAlertContent("Product deleted successfully");
            setAlertConfirmText("OK");
            setAlertStatus("success");
            setAlertCancelBottom(false);
            setAlertOnConfirm(() => () => setOpenAlert(false));
            setOpenAlert(true);
            loadProducts();
        }
        else{
            setAlertTitle("Delete Product");
            setAlertContent("Product deletion failed");
            setAlertConfirmText("OK");
            setAlertStatus("error");
            setAlertCancelBottom(false);
            setAlertOnConfirm(() => () => setOpenAlert(false));
            setOpenAlert(true);
        }
    }

    const handleDelete = async (id:string) => {
        console.log("Delete product with idaaa:", id);
        setAlertTitle("Delete Product");
        setAlertContent("Are you sure you want to delete this product?");
        setAlertConfirmText("Delete");
        setAlertStatus("warning");
        setAlertOnConfirm(() => () => deleteProduct(id));
        setAlertCancelBottom(true);
        setAlertOnCancel(() => () => setOpenAlert(false));
        setOpenAlert(true);
    };

    // const openEditModal = (product) => {
    //     // Logic to open a modal and edit the product
    //     console.log('Editing product:', product);
    // };

  return (
    <>
        <AlertDialog 
            open={openAlert} 
            setOpen={setOpenAlert} 
            title={alertTitle} 
            content={alertContent} 
            status={alertStatus} 
            onConfirm={alertOnConfirm} 
            onCanceled={alertOnCancel}
            confirmText={alertConfirmText} 
            cancelBottom={alertCancelBottom}
        />
        <DataTable columns={columns} data={products} />
    </>
    
  );
}

export default ProductTable;