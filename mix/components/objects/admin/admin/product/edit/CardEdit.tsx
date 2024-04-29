"use client"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"

import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { getAuthor,getCategory,getPublisher,getProductById,putProductById} from "@/apis/services/productService"
import { ChangeEvent, useEffect, useState } from "react"
import { AuthProvider } from "@/utils/clientAuthProvider"
import { ProductModel } from "@/models/dto/product"
import { set } from "date-fns"
import { Value } from "@radix-ui/react-select"
import Image from "next/image"
import { Button } from '@/components/ui/button';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { ProductUpdate } from "@/models/dto/product"
import { useRouter } from "next/navigation"


export default function CardEdit({ id }: { id: string }) {
    const token = AuthProvider.getAccessToken()
    const router = useRouter();

    const [author, setAuthor] = useState([{id:"",name:"",description:""}])
    const [category, setCategory] = useState([{id:"",name:""}])
    const [publisher, setPublisher] = useState([{id:"",name:""}])
    const [products, setProducts] = useState(
        {
            id: "",
            name: "",
            author:{id:"",name:"asd",description:""},
            category:{id:"",name:"asd"},
            publisher:{id:"",name:""},
            description: "",
            price: "",
            stock: "",
            img_url: "",
            status: "",
            created_at: "",
            updated_at: "",
        }
    )

    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertOnCancel, setAlertOnCancel] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');
    const [alertCancelBottom, setAlertCancelBottom] = useState(false);

    const [selectedPublisher, setSelectedPublisher] = useState('');
    const [selectedCategory, setSelectedCategory] = useState('');
    const [selectedAuthor, setSelectedAuthor] = useState('');
    //data product
    

    useEffect(() => {
        loadAuthor();
        loadCategory();
        loadPublisher();
        loadProducts();

    }, []);


    const loadAuthor = async () => {
        const fetchedAuthor = await getAuthor(token||'');
        setAuthor(fetchedAuthor.data.data);
    };
    const loadCategory = async () => {
        const fetchedCategory = await getCategory(token||'');
        setCategory(fetchedCategory.data.data);
    };
    const loadPublisher = async () => {
        const fetchedPublisher = await getPublisher(token||'');
        setPublisher(fetchedPublisher.data.data);
    };
    const loadProducts = async () => {
        const fetchedProducts = await getProductById(id,token||'');
        const product = fetchedProducts.data.data

        setProducts(product);
    };

    const handleNameInputChange = (value: string) => {
        const newProducts = { ...products };
        newProducts.name = value;
        setProducts(newProducts);
    };
    const handlePriceInputChange = (value: string) => {
        const newProducts = { ...products };
        newProducts.price = value;
        setProducts(newProducts);
    }
    const handleDescriptionInputChange = (value: string) => {
        const newProducts = { ...products };
        newProducts.description = value;
        setProducts(newProducts);
    }
    const handleStockInputChange = (value: string) => {
        const newProducts = { ...products };
        newProducts.stock = value;
        setProducts(newProducts);
    }
    


    const handlePublisherChange = (event: ChangeEvent<HTMLSelectElement>) => {
        const value = event.target.value;
        setSelectedPublisher(value);
    };

    const handleAuthorChange = (event: { target: { value: any } }) => {
        const value = event.target.value;
        setSelectedAuthor(value);
    };

    const handleCategoryChange = (event: { target: { value: any } }) => {
        const value = event.target.value;
        setSelectedCategory(value);
    };

    const handleSubmit = async (event:any) => {
        
        event.preventDefault();
        var author_id = products.author.id;
        var category_id = products.category.id;
        var publisher_id = products.publisher.id;
        

        if(selectedAuthor !== ''){
            author_id = selectedAuthor;
        }
        if(selectedCategory !== ''){
            category_id = selectedCategory;
        }
        if(selectedPublisher !== ''){
            publisher_id = selectedPublisher;
        }

        const data : ProductUpdate = {
            author_id: parseInt(author_id),
            category_id: parseInt(category_id),
            publisher_id: parseInt(publisher_id),
            name: products.name,
            description: products.description,
            price: parseFloat(products.price),
            stock: parseInt(products.stock),
            img: products.img_url,
        }
        console.log(data);
        // console.log(data);
        putProductById(id,data,token||'').then((res) => {
            if (res) {
                if (res.status === 200) {
                    console.log("success");
                    setAlertTitle("Update Success");
                    setAlertContent("Product updated successfully");
                    setAlertConfirmText("OK");
                    setAlertStatus("success");
                    setAlertCancelBottom(false);
                    setAlertOnConfirm(() => () => router.push('/dashboard/product'));
                    setOpenAlert(true);
                } else {
                    console.log("failed");
                    setAlertTitle("Update Failed");
                    setAlertContent("Product update failed");
                    setAlertConfirmText("OK");
                    setAlertStatus("error");
                    setAlertCancelBottom(false);
                    setAlertOnConfirm(() => () => setOpenAlert(false));
                    setOpenAlert(true);
                }

            } else{
                console.log('error');
            }
        })

    const backHandle = async (event:any) => {
            router.push('/dashboard/product');
    }

    }
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
            <Card>
            <CardHeader>
                <CardTitle>Edit Product</CardTitle>
                <CardDescription>Edit the product details</CardDescription>
            </CardHeader>
            <CardContent>
                <form onSubmit={handleSubmit}>
                    <div className="grid grid-rows-5 grid-flow-col gap-4">

                    
                    <div className="flex row-span-3 col-span-1 ">
                        <img src={products.img_url} alt="Product"  className="rounded-lg w-[300px] h-[150px] sm:h-[180px] md:h-[220px] 2xl:h-[260px] hover:scale-110 transition-transform duration-200" />
                    </div>
                    <div className="flex flex-col col-span-3">
                        <Label htmlFor="description">Description</Label>
                        <textarea id="description" value={products.description} className="mt-1 h-24"></textarea>
                    </div>
                    <div className="flex flex-col col-span-2">

                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="name">Name</Label>
                        <Input id="name" placeholder="Name" value={products.name} onChange={(e) => handleNameInputChange(e.target.value)} className="mt-1" />
                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="price">Price</Label>
                        <Input id="price" placeholder="Price"value={products.price} onChange={(e) => handlePriceInputChange(e.target.value)} className="mt-1" />
                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="publisher">Publisher</Label>
                        <select id="publisher" value={selectedPublisher} onChange={handlePublisherChange} className="mt-1 form-select block w-full">
                            {publisher.map((item) => (
                                <option key={item.id} value={item.id}>
                                    {item.name}
                                </option>
                            ))}
                        </select>
                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="stock">Stock</Label>
                        <Input id="stock" placeholder="Stock"value={products.stock} onChange={(e) => handleStockInputChange(e.target.value)}className="mt-1" />
                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="author">Author</Label>
                        <select id="author" value={selectedAuthor} onChange={handleAuthorChange} className="mt-1 form-select block w-full">
                            {author.map((item) => (
                                <option key={item.id} value={item.id}>
                                    {item.name}
                                </option>
                            ))}
                        </select>
                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="category">Category</Label>
                        <select id="category" value={selectedCategory} onChange={handleCategoryChange} className="mt-1 form-select block w-full">
                            {category.map((item) => (
                                <option key={item.id} value={item.id}>
                                    {item.name}
                                </option>
                            ))}
                        </select>
                    </div>
                </div>
                <div className="flex justify-end col-span-3 space-x-2 mt-4">
                    <Button onClick={()=>router.push('/dashboard/product')}  className="bg-sec ring-2 px-8 rounded-xl text-white">
                        Back
                    </Button>
                    <Button type="submit" className="bg-primary ring-2 px-8 rounded-xl text-white">
                        Save
                    </Button>
                </div>
                </form>
            </CardContent>
        </Card>
        </>
    )
}