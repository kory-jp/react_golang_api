import { useCallback, VFC } from "react"
import { Box, Flex, Link, Text } from "@chakra-ui/react"
import { useHistory } from "react-router"

export const Header: VFC = () => {
  const history = useHistory()
  const onClickTop = useCallback(()=> history.push("/"),[history])
  const onClickLogin = useCallback(()=> history.push("/login"),[history])
  const onClickSignup = useCallback(()=> history.push("/signup"),[history])
  return(
    <Flex bg="telegram.500" justifyContent="space-between">
      <Box ml="16">
        <Text fontSize="5xl" color="white">
          TodoApp
        </Text>
      </Box>
      <Flex mt="auto" mb="auto" mr="16">
        <Link
          mr="4"
          fontWeight="bold"
          color="white"
          onClick={onClickTop}
        >
          トップページ
        </Link>
        <Link 
          mr="4"
          fontWeight="bold"
          color="white"
          onClick={onClickLogin}
        >
          ログイン
        </Link>
        <Link 
          fontWeight="bold"
          color="white"
          onClick={onClickSignup}
        >
          新規登録
        </Link>
      </Flex>
    </Flex>
  )
}

export default Header;