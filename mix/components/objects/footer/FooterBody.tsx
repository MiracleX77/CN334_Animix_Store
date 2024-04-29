import Middle from '@/components/layouts/Middle'
import Link from 'next/link';


export default function FooterBody() {


    return (
        <>

            <Middle Y className="relative w-full min-h-[200px] h-max py-4 bg-card text-md text-foreground">
                <div className='w-full h-full grid grid-cols-1 md:grid-cols-6 gap-8'>
                    <div className='col-span-full md:col-span-2 h-full'>
                        <h2 className="text-2xl font-bold">
                            Florist
                        </h2>
                        <p className='text-muted-foreground flex flex-col justify-center'>
                            <span className='text-lg'>
                                Grand siam orchid Co., Ltd.
                            </span>
                            <span>
                                551/60 moo.12, Salathammasop,Thaweewatthana, Bangkok 10170 Thailand.
                            </span>
                            <span>
                                Tel : 66 34 333475 , 66 34 333459
                            </span>
                            <span>
                                Fax : 66-34 333468 ,
                            </span>
                            <span>
                                Email : &nbsp;
                                <a href="mailto: gs_orchid@hotmail.com" className='underline'>
                                    gs_orchid@hotmail.com
                                </a>
                            </span>
                            <span>
                                Contact : Ms. Pankamon Lergsri
                            </span>
                            <span>
                                Mobile phone : 66 81 6238420
                            </span>
                            <span>
                                Email : &nbsp;
                                <a href="mailto: gs_orchid@hotmail.com" className='underline'>
                                    gs_orchid@hotmail.com
                                </a>
                            </span>
                        </p>
                    </div>
                    <div className='col-span-full md:col-span-1 h-full'>
                        <h3 className='text-lg text-muted-foreground'>
                            Links
                        </h3>
                        <div className='w-full h-full grid grid-rows-6 gap-4'>
                            <div></div>
                            <Link href='/'>Home</Link>
                            <Link href='/'>Product</Link>
                            <Link href='/'>About</Link>
                            <Link href='/'>Contact</Link>
                        </div>
                    </div>
                    <div className='col-span-full md:col-span-1 h-full'>
                        <h3 className='text-lg text-muted-foreground'>
                            Help
                        </h3>
                        <div className='w-full h-full grid grid-rows-6 gap-4'>
                            <div></div>
                            <Link href='/'>Support</Link>
                        </div>
                    </div>
                </div>

            </Middle>
            <div className="w-full h-0.5 bg-muted-foreground opacity-20"></div>
        </>
    );
}
