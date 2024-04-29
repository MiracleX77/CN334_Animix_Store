'use client'
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {Login,Register} from "@/models/dto/auth"
import {
    Tabs,
    TabsContent,
    TabsList,
    TabsTrigger,
} from "@/components/ui/tabs"
import { useState } from "react"
import { postLogin,postRegister } from "@/apis/services/authService"
import { AuthProvider } from "@/utils/clientAuthProvider"
import { useRouter } from "next/navigation"
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { set } from "date-fns"


export default function Auth() {

    const router = useRouter();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [first_name, setFirstName] = useState('');
    const [last_name, setLastName] = useState('');

    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');


    const handleUsernameChange = (event:any) => setUsername(event.target.value);
    const handlePasswordChange = (event:any) => setPassword(event.target.value);
    const handleEmailChange = (event:any) => setEmail(event.target.value);
    const handleFirstNameChange = (event:any) => setFirstName(event.target.value);
    const handleLastNameChange = (event:any) => setLastName(event.target.value);


    const handleLoginSubmit = async (event:any) => {
        event.preventDefault();
        const data : Login = {
            username: username,
            password: password
        }
        postLogin(data).then((res) => {
            console.log(res);
            if (res) {
                if (res.status === 200) {
                    console.log("success");

                    AuthProvider.login(res.data.data.token, res.data.data.role);
                    router.push('/');

                    
                    //redirect to home
                } else {
                    setAlertTitle("Login failed");
                    setAlertContent("Username or password is incorrect");
                    setAlertConfirmText("OK");
                    setAlertStatus("error");
                    setAlertOnConfirm(() => () => setOpenAlert(false));
                    setOpenAlert(true);
                    setUsername('');
                    setPassword('');

                }
            } else{
                console.log('error');
            }
        })
    }
    const handleRegisterSubmit = async (event:any) => {
        event.preventDefault();
        const data : Register = {
            first_name: first_name,
            last_name: last_name,
            email: email,
            username: username,
            password: password
        }
        console.log(data);
        postRegister(data).then((res) => {
            if (res) {
                if (res.status === 200) {
                    console.log("success");
                    //redirect to 
                    setAlertTitle("Register success");
                    setAlertContent("Register success");
                    setAlertConfirmText("OK");
                    setAlertStatus("success");
                    setAlertOnConfirm(() => () => {
                        setOpenAlert(false);
                        router.push('/auth');
                    });
                    setOpenAlert(true);

                    
                } else {
                    if (res.message === "username already exist") {
                        setAlertTitle("Register failed");
                        setAlertContent("Username already exist");
                        setAlertConfirmText("OK");
                        setAlertStatus("error");
                        setAlertOnConfirm(() => () => setOpenAlert(false));
                        setOpenAlert(true);
                        setUsername('');
                        setPassword('');
                        setEmail('');
                        setFirstName('');
                        setLastName('');
                    } else{
                        setAlertTitle("Register failed");
                        setAlertContent("Something went wrong");
                        setAlertConfirmText("OK");
                        setAlertStatus("error");
                        setAlertOnConfirm(() => () => setOpenAlert(false));
                        setOpenAlert(true);
                        setUsername('');
                        setPassword('');
                        setEmail('');
                        setFirstName('');
                        setLastName('');
                    }
                }

            } else{
                console.log('error');
            }
        })

    }



    return (
        <>
            <AlertDialog open={openAlert} setOpen={setOpenAlert} title={alertTitle} content={alertContent} status={alertStatus} onConfirm={alertOnConfirm} confirmText={alertConfirmText} cancelBottom={false}/>
            <Tabs defaultValue="login" className="w-[400px]">
                <TabsList className="grid w-full grid-cols-2">
                    <TabsTrigger value="login">Login</TabsTrigger>
                    <TabsTrigger value="register">Register</TabsTrigger>
                </TabsList>
                <TabsContent value="login">
                    <Card>
                        <form onSubmit={handleLoginSubmit}>
                            <CardHeader>
                                <CardTitle>Login</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-2">
                                <div className="space-y-1">
                                    <Label htmlFor="username">Username</Label>
                                    <Input id="username" value={username} required placeholder="Enter your username" onChange={handleUsernameChange} />
                                </div>
                                <div className="space-y-1">
                                    <Label htmlFor="password">Password</Label>
                                    <Input id="password" type="password" value={password} required placeholder="Enter your password" onChange={handlePasswordChange} />
                                </div>
                            </CardContent>
                            <CardFooter>
                                <Button  type="submit" >Login</Button>
                            </CardFooter>
                        </form>
                    </Card>
                </TabsContent>
                <TabsContent value="register">
                    <Card>
                        <form onSubmit={handleRegisterSubmit}>
                            <CardHeader>
                                <CardTitle>Register</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-1">
                                <div className="flex flex-wrap -mx-2">
                                    <div className="px-2 w-1/2">
                                        <Label htmlFor="first_name">First name</Label>
                                        <Input id="first_name" value={first_name} required placeholder="Enter your First name" onChange={handleFirstNameChange}/>
                                    </div>
                                    <div className="px-2 w-1/2">
                                        <Label htmlFor="last_name">Last name</Label>
                                        <Input id="last_name"  value={last_name} required placeholder="Enter your Last name" onChange={handleLastNameChange} />
                                    </div>
                                </div>
                                <div className="space-y-1">
                                    <Label htmlFor="email">Email</Label>
                                    <Input id="email" type="email" value={email} required placeholder="Enter your email" onChange={handleEmailChange}/>
                                </div>
                                <div className="space-y-1">
                                    <Label htmlFor="username">Username</Label>
                                    <Input id="username" type="username" value={username} required placeholder="Enter your username" onChange={handleUsernameChange}/>
                                </div>
                                <div className="space-y-1">
                                    <Label htmlFor="password">Password</Label>
                                    <Input id="password" type="password" value={password} required placeholder="Enter your password" onChange={handlePasswordChange}/>
                                </div>
                            </CardContent>
                            <CardFooter>
                                <Button type="submit" >Save password</Button>
                            </CardFooter>
                        </form>
                    </Card>
                </TabsContent>
            </Tabs>

        </>
    )
}