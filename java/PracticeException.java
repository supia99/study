class PracticeException {

  static void main(String args[]) {

    try {
      throw new Exception() ;
    }catch(Exception ex) {
      System.out.println("catch Exception");
      throw ex;
    }finally {
      try {
        throw new Exception();
      } catch(Exception ex) {
        ex.stackTrace();
      }
    }
  }
}
