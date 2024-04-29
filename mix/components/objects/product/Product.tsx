import { Button } from '@/components/ui/button';
import Quantity from '@/components/ui/quantity';
import Image from 'next/image';
import { useState } from 'react';

export default function ProductComponent({ id }: { id: string }) {

    const [quantity, setQuantity] = useState(1)

    const handleSetQuantity = (quantity: number) => {
        if (quantity < 1 || quantity === null || quantity === undefined || isNaN(quantity)) {
            setQuantity(1)
        } else {
            setQuantity(quantity)
        }
    }

    return (
        <>
            <div className="w-full h-full grid grid-cols-1 md:grid-cols-2 gap-8">
                <div>
                    <div>
                        <Image src="https://via.placeholder.com/1440" alt="product" width={1440} height={1440} layout='cover' className='w-full h-[300px] sm:h-[500px] md:h-[600px] 2xl:h-[900px] rounded-2xl' />
                    </div>
                    <div className='grid grid-cols-2 py-4 gap-4'>
                        <Image src="https://via.placeholder.com/1440" alt="product" width={1440} height={1440} layout='cover' className='w-full h-[400px] rounded-2xl' />
                        <Image src="https://via.placeholder.com/1440" alt="product" width={1440} height={1440} layout='cover' className='w-full h-[400px] rounded-2xl' />
                    </div>
                </div>
                <div>
                    <h2 className='text-4xl font-bold'>Product Name</h2>
                    <p className='text-sm text-secondary-foreground'>วันวางจำหน่าย : 14-11-2018</p>
                    <p className='text-lg'>ราคา : ฿ 1000</p>
                    <div className='flex items-center space-x-4'><p>จำนวน :</p><Quantity quantity={1} setQuantity={() => { }} /></div>
                    <Button className="bg-secondary ring-2 px-4 rounded-xl text-white">
                        Add
                    </Button>
                    <div>
                        <h2 className='text-2xl font-bold'>Description</h2>
                        <p>
                            สาธารณรัฐซันแมกโนเลียถูก ‘จักรวรรดิ’ ซึ่งเป็นดินแดนข้างเคียงส่งจักรกลรบไร้พลขับ‘ลีเจี้ยน’ เข้ามารุกรานไม่เว้นแต่ละวัน สาธารณรัฐจึงคิดค้นอาวุธรูปแบบเดียวกันขึ้นเพื่อการตอบโต้และในที่สุดก็สามารถขับไล่ภัยคุกคามสิ้นไปโดยไม่มีผู้ใดต้องสละชีพใช่―เพียงเปลือกนอกความจริงแล้วหาใช่ปราศจากการหลั่งเลือดเนื้อ‘เขต 86 ไร้ซึ่งตัวตน’ ตั้งอยู่นอก85เขตของสาธารณรัฐทุกวันคืน เหล่าหนุ่มสาวอันถูกตีตรา ‘เอทตี้ซิกซ์’ ได้เข้าต่อสู้ในฐานะ ‘จักรกลมีพลขับ’ ตลอดมา―เด็กหนุ่มนามว่าชิน ผู้นำเหล่าหนุ่มสาวเข้าสู่ดินแดนแห่งความตายเด็กสาวนามเรน่า ผู้เป็นแฮนด์เลอร์-ผู้บังคับบัญชาการออกคำสั่งพวกเขาผ่านการสื่อสารพิเศษจากแนวหลังอันห่างไกล|เรื่องราวของทั้งสอง...การต่อสู้ซึ่งเป็นโศกนาฏกรรมแสนทารุณและจากลานับนิรันดร์เริ่มต้นขึ้นแล้ว―!ผลงานชิ้นเอก เจ้าของรางวัลเกียรติยศ ‘รางวัลชนะเลิศ’ การประกวดเด็งเกคิโนเวลไพร์ซครั้งที่ 23เดินหน้าเข้าประจัญบานอย่างสมศักดิ์ศรี!
                        </p>
                    </div>
                </div>
            </div>
        </>
    )
}

