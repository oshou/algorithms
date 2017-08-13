require "test/unit"
require "minitest/autorun"
require "sorts"
@@results = {}

class TestSorts < Test::Unit::TestCase
  def setup
    num = 1000
    @list = []
    num.times { @list << rand(num) }
  end

  def time(name)
    t = Time.now
    res = yield
    @@result[name] = Time.now - t
    res
  end

  def test_insert_sort
    assert_equal(@list.sort,time(:Insert){ @list.insert_sort })
  end

  def test_select_sort
    assert_equal(@list.sort,time(:Select){ @list.select_sort })
  end

  def test_bubble_sort
    assert_equal(@list.sort,time(:Bubble){ @list.bubble_sort })
  end

  def test_quick_sort
    assert_equal(@list.sort,time(:Quick){ @list.quick_sort })
  end

  def test_merge_sort
    assert_equal(@list.sort,time(:Merge){ @list.merge_sort})
  end

  def test_sort
    assert_equal(@list.sort,time(:Sort){ @list.sort})
  end
end

END{
  END{
    res = @@result.map{|k,v|[k,v,(v/@@result[:Sort]).to_i]}.sort_by{|e| e[2]}
    res.each{ |k,v,n| puts "#{k}\t=>\t#{v} (#{n})"}
  }
}
