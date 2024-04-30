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
import { getAuthor,getCategory,getPublisher,postProduct} from "@/apis/services/productService"
import { ChangeEvent, useEffect, useState } from "react"
import { AuthProvider } from "@/utils/clientAuthProvider"
import { ProductModel } from "@/models/dto/product"
import { set } from "date-fns"
import { Value } from "@radix-ui/react-select"
import Image from "next/image"
import { Button } from '@/components/ui/button';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { ProductCreate } from "@/models/dto/product"
import { useRouter } from "next/navigation"


export default function CardAdd() {
    const token = AuthProvider.getAccessToken()
    const router = useRouter();

    const [author, setAuthor] = useState([{id:"",name:"",description:""}])
    const [category, setCategory] = useState([{id:"",name:""}])
    const [publisher, setPublisher] = useState([{id:"",name:""}])
    const [products, setProducts] = useState(
        {
            name: "",
            author_id: 0,
            category: 0,
            publisher: 0,
            description: "",
            price: 0,
            stock: 0,
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
    //file image
    const [image, setImage] = useState<File | null>(null);
    const [imagePreview, setImagePreview] = useState<string | null>(null);
    
    useEffect(() => {
        loadAuthor();
        loadCategory();
        loadPublisher();
    }, []);


    const loadAuthor = async () => {
        const fetchedAuthor = await getAuthor(token||'');
        setSelectedAuthor(fetchedAuthor.data.data[0].id);
        setAuthor(fetchedAuthor.data.data);
    };
    const loadCategory = async () => {
        const fetchedCategory = await getCategory(token||'');
        setSelectedCategory(fetchedCategory.data.data[0].id);
        setCategory(fetchedCategory.data.data);
    };
    const loadPublisher = async () => {
        const fetchedPublisher = await getPublisher(token||'');
        console.log(fetchedPublisher);
        setSelectedPublisher(fetchedPublisher.data.data[0].id);
        setPublisher(fetchedPublisher.data.data);
    };

    const handleNameInputChange = (value: string) => {
        const newProducts = { ...products };
        newProducts.name = value;
        setProducts(newProducts);
    };
    const handlePriceInputChange = (value: number) => {
        const newProducts = { ...products };
        newProducts.price = value;
        setProducts(newProducts);
    }
    const handleDescriptionInputChange = (value: string) => {
        const newProducts = { ...products };
        newProducts.description = value;
        setProducts(newProducts);
    }
    const handleStockInputChange = (value: number) => {
        const newProducts = { ...products };
        newProducts.stock = value;
        setProducts(newProducts);
    }
    

    const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files && event.target.files[0];
        if (file) {
            console.log(file);
            setImage(file);
            const reader = new FileReader();
            reader.onloadend = () => {
                // When the file is read, set the image preview state
                setImagePreview(reader.result as string);
            };
            reader.readAsDataURL(file);
        }
    };


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
        
        console.log(selectedAuthor);
        console.log(selectedCategory);
        console.log(selectedPublisher);
        event.preventDefault();

        var author_id = selectedAuthor;
        var category_id = selectedCategory;
        var publisher_id = selectedPublisher;


        const data : ProductCreate = {
            author_id: parseInt(author_id),
            category_id: parseInt(category_id),
            publisher_id: parseInt(publisher_id),
            name: products.name,
            description: products.description,
            price: products.price,
            stock: products.stock,
            img: image ,
        }
        console.log(data);
        // console.log(data);
        postProduct(data,token||'').then((res) => {
            if (res) {
                if (res.status === 200) {
                    console.log("success");
                    setAlertTitle("Product Added");
                    setAlertContent("Product added successfully");
                    setAlertConfirmText("OK");
                    setAlertStatus("success");
                    setAlertCancelBottom(false);
                    setAlertOnConfirm(() => () => {
                        setOpenAlert(false);
                        router.push('/dashboard/product');
                    });
                    setOpenAlert(true);
                } else {
                    console.log("error");
                    setAlertTitle("Add Product");
                    setAlertContent("Product addition failed");
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
                <CardTitle>Add Product</CardTitle>
                <CardDescription>Edit the product details</CardDescription>
            </CardHeader>
            <CardContent>
        <form onSubmit={handleSubmit} className="grid grid-cols-4 gap-4">
            {/* File input and Image preview occupying the first column and spanning 2 rows */}
            <div className="flex flex-col col-span-1 row-span-4">
                <div className="">
                    <Label htmlFor="image">Image</Label>
                    <input 
                        id="image" 
                        type="file" 
                        accept="image/*" 
                        onChange={handleImageChange} 
                        className="mt-1"
                    />
                </div>
                <div className="flex-grow">
                    {imagePreview && (
                        <img 
                            src={imagePreview} 
                            alt="Preview" 
                            className="rounded-lg w-full h-auto max-h-[300px] object-cover mt-2" 
                        />
                    )}
                </div>
            </div>
            {/* Other input fields occupying the next 3 columns */}
            <div className="flex flex-col col-span-1">
                <Label htmlFor="name">Name</Label>
                <Input id="name" placeholder="Name..."  onChange={(e) => handleNameInputChange(e.target.value)} className="mt-1" />
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="price">Price</Label>
                <Input id="price" placeholder="Price..."  onChange={(e) => handlePriceInputChange(Number(e.target.value))} className="mt-1" />
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="stock">Stock</Label>
                <Input id="stock" placeholder="Stock..." onChange={(e) => handleStockInputChange(Number(e.target.value))} className="mt-1" />
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="author">Author</Label>
                    <select id="author" onChange={handleAuthorChange} className="mt-1 form-select block w-full">
                        {author.map((item) => (
                        <option key={item.id} value={item.id}>
                                {item.name}
                            </option>
                        ))}
                </select>
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="publisher">Publisher</Label>
                    <select id="publisher"  onChange={handlePublisherChange} className="mt-1 form-select block w-full">
                        {publisher.map((item) => (
                        <option key={item.id} value={item.id}>
                                {item.name}
                            </option>
                        ))}
                </select>
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="category">Category</Label>
                    <select id="category"  onChange={handleCategoryChange} className="mt-1 form-select block w-full">
                        {category.map((item) => (
                        <option key={item.id} value={item.id}>
                                {item.name}
                            </option>
                        ))}
                </select>
            </div>
            
            {/* ...more input fields... */}
            <div className="flex flex-col col-span-3">
                <Label htmlFor="description">Description</Label>
                <textarea id="description" onChange={(e) => handleDescriptionInputChange(e.target.value)} className="mt-1 h-24"></textarea>
            </div>
            <div className="col-span-3 flex justify-end space-x-2 mt-4">
                <Button onClick={()=>router.push('/dashboard/product')} className="bg-secondary ring-2 px-8 rounded-xl text-white">
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