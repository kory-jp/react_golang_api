import { useCallback, VFC } from "react"
import { Box, Flex, Heading, Divider, Input, Stack } from "@chakra-ui/react"
import { PrimaryButton } from "../atoms/button/PrimaryButton";
import { useHistory } from "react-router";

export const Singup: VFC = () => {
  const history = useHistory()
  const onClickIndex = useCallback(()=> history.push("/todo"),[history])
  return(
    <Flex align="center" justify="center" h="100vh">
    <Box
      bg="white"
      p="2"
      pl="6"
      pr="6"
      borderRadius="md"
      shadow="md"
      w="lg"
    >
      <Stack spacing="9" mb="8">
        <Heading textAlign="center" color="telegram.500">
          サインアップ
        </Heading>
        <Divider />
        <Input
          placeholder="お名前"
          />
        <Input
          placeholder="メールアドレス"
          />
        <Input
          placeholder="パスワード"
          />
          <PrimaryButton
            onClick={onClickIndex}
          >
            ログイン
          </PrimaryButton>
      </Stack>
    </Box>
  </Flex>
  )
}

export default Singup;