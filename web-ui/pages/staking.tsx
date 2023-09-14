import {
  Box,
  Image,
  Container,
  Flex,
  VStack,
  HStack,
  Stat,
  StatLabel,
  StatNumber,
} from '@chakra-ui/react';
import dynamic from 'next/dynamic';
import Head from 'next/head';
import { useState } from 'react';

import { Header } from '@/components';
import { SideHeader } from '@/components';
import { NetworkSelect } from '@/components';
import { StakingBox } from '@/components';
import { InfoBox } from '@/components';
import { AssetsAccordian } from '@/components';

const DynamicStakingBox = dynamic(() => Promise.resolve(StakingBox), {
  ssr: false,
});

export default function Staking() {
  const [selectedOption, setSelectedOption] = useState('Atom');
  const [selectedChainName, setSelectedChainName] = useState('cosmoshub');
  const [isModalOpen, setModalOpen] = useState(false);
  useState(null);
  return (
    <>
      <Box
        w="100vw"
        h="100vh"
        bgImage="url('https://s3-alpha-sig.figma.com/img/555d/db64/f5bf65e93a15603069e8e865d5f6d60d?Expires=1694995200&Signature=fYfmbqDdOGRYtSeEsOkavPhhkaNQK1UFFfICaUaM1k9OVEpACsoWOcK2upjRW7Tfs-pPTJBuQuvcmF9gBjosh5-Al2xTWHYzDlR~CYJNzsXcseIEnVf7H8lCdJqhZY-T0r~lmbJK5-CmbulWfOaubc-wyY3C-oM3b1RanGV1TqmPZto5bbHwf56jDYqK86HedVMXbUCOlzkeBw2R93AkmNDMOdDbKa9rIKqxil64DuQQAfIFxWm1Rc69Jc1-4K-bunsS~kfz8bSET6TIGmR15nCo~ibfISG72YYKAa7zz6XqUY6GKmmG-Yhj9XyyYb7Jy02r5axNei3DRD78SBe~6w__&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4')"
        bgSize="fit"
        bgPosition="right center"
        bgAttachment="fixed"
        bgRepeat="no-repeat"
        bgColor="#000000"
      >
        <Head>
          <title>Staking</title>
          <meta
            name="viewport"
            content="width=device-width, initial-scale=1.0"
          />
          <link rel="icon" href="/img/favicon.png" />
        </Head>
        <Header chainName={selectedChainName} />
        <SideHeader />
        <Container
          zIndex={2}
          position="relative"
          mt={-7}
          maxW="container.lg"
          maxH="80vh"
          h="80vh"
        >
          <Image
            alt={''}
            src="/img/metalmisc2.png"
            zIndex={-10}
            position="absolute"
            bottom="-10"
            left="-10"
            boxSize="120px"
          />
          <Flex zIndex={3} direction="column" h="100%">
            {/* Dropdown and Statistic */}
            <Box w="50%">
              <HStack justifyContent="space-between" w="100%">
                <NetworkSelect
                  selectedOption={selectedOption}
                  setSelectedOption={setSelectedOption}
                  selectedChainName={selectedChainName}
                  setSelectedChainName={setSelectedChainName}
                />
                <VStack p={1} borderRadius="10px" alignItems="flex-end">
                  <Stat color="complimentary.900">
                    <StatLabel>APY</StatLabel>
                    <StatNumber>35%</StatNumber>
                  </Stat>
                </VStack>
              </HStack>
            </Box>

            {/* Content Boxes */}
            <Flex h="100%">
              {/* Staking Box*/}
              <DynamicStakingBox
                selectedOption={selectedOption}
                isModalOpen={isModalOpen}
                setModalOpen={setModalOpen}
                selectedChainName={selectedChainName}
              />

              <Box w="10px" />

              {/* Right Box */}
              <Flex flex="1" direction="column">
                {/* Top Half (2/3) */}
                <InfoBox selectedOption={selectedOption} />

                <Box h="10px" />
                {/* Bottom Half (1/3) */}
                <AssetsAccordian selectedOption={selectedOption} />
              </Flex>
            </Flex>
          </Flex>
        </Container>
      </Box>
    </>
  );
}
